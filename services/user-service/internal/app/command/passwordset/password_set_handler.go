package passwordset

import (
	"context"
	"fmt"

	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/service"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/util"
)

type PasswordSetCommandHandler struct {
	config       *config.Config
	userRepo     repo.UserRepo
	tokenRepo    repo.TokenRepo
	emailService client.EmailServicePort
}

func NewPasswordSetCommandHandler(
	config *config.Config,
	userRepo repo.UserRepo,
	tokenRepo repo.TokenRepo,
	emailService client.EmailServicePort,
) *PasswordSetCommandHandler {
	return &PasswordSetCommandHandler{
		config:       config,
		userRepo:     userRepo,
		tokenRepo:    tokenRepo,
		emailService: emailService,
	}
}

func (h *PasswordSetCommandHandler) Handle(ctx context.Context, command *PasswordSetCommand) (*PasswordSetCommandResponse, error) {

	user, err := h.userRepo.FindByID(ctx, command.UserID)
	if err != nil {
		return nil, err
	}

	if !user.GetIsVerified() {
		return nil, fmt.Errorf("please verify your email first")
	}

	setTokenData, err := h.tokenRepo.Find(ctx, command.UserID, command.SetToken)
	if err != nil || setTokenData == nil {
		return nil, err
	}

	hashPassword, err := util.HashPassword(command.Password)
	if err != nil {
		return nil, err
	}

	user.SetPassword(hashPassword)

	_, err = h.userRepo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	err = h.tokenRepo.Delete(ctx, command.UserID, command.SetToken)
	if err != nil {
		return nil, err
	}

	go func() {

		contentData := service.NewEmailContentData(
			"Password Set Confirmation",
			"Password Set Confirmation",
			user.GetFirstName(),
			user.GetLastName(),
			fmt.Sprintf("%s/api/v1", h.config.FRONTEND_URL),
		)

		err = h.emailService.SendEmail(
			h.config.SENDER_EMAIL,
			[]string{user.GetEmail()}, nil, nil,
			"Congratulations! Your password has been set",
			"password_set_success",
			contentData,
		)
		if err != nil {
			fmt.Printf("Error sending password set confirmation email: %v", err)
		}
	}()

	return &PasswordSetCommandResponse{
		Message: "Password set successfully",
	}, nil
}
