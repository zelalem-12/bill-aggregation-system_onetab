package passwordchange

import (
	"context"

	"github.com/zelalem-12/onetab/internal/app/repo"
	"github.com/zelalem-12/onetab/internal/util"
)

type PasswordChangeCommandHandler struct {
	userRepo  repo.UserRepo
	tokenRepo repo.TokenRepo
}

func NewPasswordChangeCommandHandler(
	UserRepo repo.UserRepo,
	tokenRepo repo.TokenRepo,
) *PasswordChangeCommandHandler {
	return &PasswordChangeCommandHandler{
		userRepo:  UserRepo,
		tokenRepo: tokenRepo,
	}
}

func (h *PasswordChangeCommandHandler) Handle(ctx context.Context, command *PasswordChangeCommand) (*PasswordChangeCommandResponse, error) {

	user, err := h.userRepo.FindByID(ctx, command.UserID)
	if err != nil {
		return nil, err
	}

	if !util.ComparePasswords(command.OldPassword, user.GetPassword()) {
		return nil, util.CreateError("Old password is incorrect")
	}

	hashedPassword, err := util.HashPassword(command.NewPassword)
	if err != nil {
		return nil, err
	}

	user.SetPassword(hashedPassword)

	_, err = h.userRepo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	// Revoke User Refresh Tokens
	err = h.tokenRepo.DeleteByUserID(ctx, command.UserID)
	if err != nil {
		return nil, err
	}

	return &PasswordChangeCommandResponse{
		Message: "Password changed successfully. Please log in again.",
	}, nil
}
