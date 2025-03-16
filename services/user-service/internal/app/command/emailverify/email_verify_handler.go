package emailverify

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zelalem-12/onetab/internal/app/client"
	"github.com/zelalem-12/onetab/internal/app/repo"
	"github.com/zelalem-12/onetab/internal/app/service"
	"github.com/zelalem-12/onetab/internal/domain"
	"github.com/zelalem-12/onetab/internal/infrastructure/config"
	"github.com/zelalem-12/onetab/internal/util"
)

type EmailVerifyCommandHandler struct {
	config       *config.Config
	userRepo     repo.UserRepo
	tokenRepo    repo.TokenRepo
	emailService client.EmailServicePort
}

func NewEmailVerifyCommandHandler(
	config *config.Config,
	userRepo repo.UserRepo,
	tokenRepo repo.TokenRepo,
	emailService client.EmailServicePort,
) *EmailVerifyCommandHandler {
	return &EmailVerifyCommandHandler{
		config:       config,
		userRepo:     userRepo,
		tokenRepo:    tokenRepo,
		emailService: emailService,
	}
}

func (h *EmailVerifyCommandHandler) Handle(ctx context.Context, command *EmailVerifyCommand) (*EmailVerifyCommandResponse, error) {

	user, err := h.userRepo.FindByID(ctx, command.UserID)
	if err != nil {
		return nil, err
	}

	if user.GetIsVerified() {
		return nil, util.CreateError("user already verified")
	}

	userID, err := uuid.Parse(user.GetID())
	if err != nil {
		return nil, err
	}

	tokenData, err := h.tokenRepo.Find(ctx, command.UserID, command.Token)
	if err != nil || tokenData == nil {
		return nil, util.CreateError("invalid token")
	}

	user.SetIsVerified(true)

	_, err = h.userRepo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	err = h.tokenRepo.Delete(ctx, command.UserID, command.Token)
	if err != nil {
		return nil, err
	}

	go func() {
		passwordSetToken, err := util.GenerateNonAccessToken(h.config.ACCESS_TOKEN_KEY, userID, h.config.ACCESS_TOKEN_EXPIRY)
		if err != nil {
			fmt.Printf("Error generating password set token :%v", err)
		}

		_, err = h.tokenRepo.Save(ctx, domain.NewToken(passwordSetToken, user.GetID()))
		if err != nil {
			fmt.Printf("Error saving password set token :%v", err)
		}

		contentData := service.NewEmailContentData(
			"Congratulations! Your email has been verified",
			"Congratulations! Your email has been verified",
			user.GetFirstName(),
			user.GetLastName(),
			fmt.Sprintf("%s/reset-password?token=%s", h.config.FRONTEND_URL, passwordSetToken),
		)

		err = h.emailService.SendEmail(
			h.config.SENDER_EMAIL,
			[]string{user.GetEmail()}, nil, nil,
			"Congratulations! Your email has been verified",
			"email_verification_success",
			contentData,
		)
		if err != nil {
			fmt.Printf("Error sending email verificaiton email: %v", err)
		}
	}()

	return &EmailVerifyCommandResponse{
		Message: "Email Verified Successfully, Please check your email to set your password",
	}, nil
}
