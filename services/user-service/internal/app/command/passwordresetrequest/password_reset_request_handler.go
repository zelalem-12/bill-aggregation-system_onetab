package passwordresetrequest

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/service"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/domain"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/util"
)

type PasswordResetRequestCommandHandler struct {
	config       *config.Config
	userRepo     repo.UserRepo
	tokenRepo    repo.TokenRepo
	emailService client.EmailServicePort
}

func NewPasswordResetRequestCommandHandler(
	config *config.Config,
	userRepo repo.UserRepo,
	tokenRepo repo.TokenRepo,
	emailService client.EmailServicePort,
) *PasswordResetRequestCommandHandler {
	return &PasswordResetRequestCommandHandler{
		config:       config,
		userRepo:     userRepo,
		tokenRepo:    tokenRepo,
		emailService: emailService,
	}
}

func (h *PasswordResetRequestCommandHandler) Handle(ctx context.Context, command *PasswordResetRequestCommand) (*PasswordResetRequestCommandResponse, error) {

	user, err := h.userRepo.FindByEmail(ctx, command.Email)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(user.GetID())
	if err != nil {
		return nil, err
	}

	if token, err := h.tokenRepo.FindByUserID(ctx, userID); err == nil && token != nil {
		err = h.tokenRepo.DeleteByUserID(ctx, userID)
		if err != nil {
			return nil, err
		}
	}

	resetToken, err := util.GenerateNonAccessToken(h.config.ACCESS_TOKEN_KEY, userID, h.config.ACCESS_TOKEN_EXPIRY)
	if err != nil {
		return nil, err
	}

	//TODO use  cache only the reset token instead of saving it to the database

	_, err = h.tokenRepo.Save(ctx, domain.NewToken(resetToken, user.GetID()))
	if err != nil {
		return nil, err
	}
	go func() {

		contentData := service.NewEmailContentData(
			"Password Reset Request",
			"Password Reset Request",
			user.GetFirstName(),
			user.GetLastName(),
			fmt.Sprintf("%s/reset-password?token=%s", h.config.FRONTEND_URL, resetToken),
		)

		err = h.emailService.SendEmail(
			h.config.SENDER_EMAIL,
			[]string{user.GetEmail()}, nil, nil,
			"Reset Your Password ",
			"password_reset_request",
			contentData,
		)
		if err != nil {
			fmt.Printf("Error sending password reset email: %v", err)
		}
	}()

	return &PasswordResetRequestCommandResponse{
		Message: "Password reset request sent successfully. Please check your email.",
	}, nil
}
