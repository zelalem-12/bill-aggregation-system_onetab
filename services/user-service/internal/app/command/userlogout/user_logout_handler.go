package userlogout

import (
	"context"

	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/util"
)

type UserLogoutCommandHandler struct {
	userRepo  repo.UserRepo
	tokenRepo repo.TokenRepo
}

func NewUserLogoutCommandHandler(
	userRepo repo.UserRepo,
	tokenRepo repo.TokenRepo,
) *UserLogoutCommandHandler {
	return &UserLogoutCommandHandler{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (h *UserLogoutCommandHandler) Handle(ctx context.Context, command *UserLogoutCommand) (*UserLogoutCommandResponse, error) {

	user, err := h.userRepo.FindByID(ctx, command.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, util.CreateError("user not found")
	}

	if token, err := h.tokenRepo.FindByUserID(ctx, command.UserID); err == nil && token != nil {
		err = h.tokenRepo.DeleteByUserID(ctx, command.UserID)
		if err != nil {
			return nil, err
		}
	}

	return &UserLogoutCommandResponse{
		Message: "User logged out successfully",
	}, nil
}
