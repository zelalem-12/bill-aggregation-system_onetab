package passwordreset

import (
	"context"
	"fmt"

	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/service"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/util"
)

type PasswordResetCommandHandler struct {
	config       *config.Config
	userRepo     repo.UserRepo
	tokenRepo    repo.TokenRepo
	emailService client.EmailServicePort
}

func NewPasswordResetCommandHandler(
	config *config.Config,
	userRepo repo.UserRepo,
	tokenRepo repo.TokenRepo,
	emailService client.EmailServicePort,
) *PasswordResetCommandHandler {
	return &PasswordResetCommandHandler{
		config:       config,
		userRepo:     userRepo,
		tokenRepo:    tokenRepo,
		emailService: emailService,
	}
}

func (h *PasswordResetCommandHandler) Handle(ctx context.Context, command *PasswordResetCommand) (*PasswordResetCommandResponse, error) {

	user, err := h.userRepo.FindByID(ctx, command.UserID)
	if err != nil {
		return nil, err
	}

	if !user.GetIsVerified() {
		return nil, fmt.Errorf("please verify your email first")
	}

	resetTokenData, err := h.tokenRepo.Find(ctx, command.UserID, command.ResetToken)
	if err != nil || resetTokenData == nil {
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

	err = h.tokenRepo.Delete(ctx, command.UserID, command.ResetToken)
	if err != nil {
		return nil, err
	}

	go func() {

		contentData := service.NewEmailContentData(
			"Password Reset Confirmation",
			"Password Reset Confirmation",
			user.GetFirstName(),
			user.GetLastName(),
			fmt.Sprintf("%s/api/v1", h.config.FRONTEND_URL),
		)

		err = h.emailService.SendEmail(
			h.config.SENDER_EMAIL,
			[]string{user.GetEmail()}, nil, nil,
			"Congratulations! Your password has been reset",
			"password_reset_success",
			contentData,
		)
		if err != nil {
			fmt.Printf("Error sending confirmation email: %v", err)
		}
	}()

	return &PasswordResetCommandResponse{
		Message: "Password reset successfully",
	}, nil
}
