package currentuserdelete

import (
	"context"
	"errors"

	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
)

type DeleteCurrentUserHandler struct {
	userRepo repo.UserRepo
}

func NewDeleteCurrentUserHandler(userRepo repo.UserRepo) *DeleteCurrentUserHandler {
	return &DeleteCurrentUserHandler{userRepo: userRepo}
}

func (h *DeleteCurrentUserHandler) Handle(ctx context.Context, command *CurrentUserDeleteCommand) (*CurrentUserDeleteCommandResponse, error) {

	if err := command.Validate(); err != nil {
		return nil, err
	}

	user, err := h.userRepo.FindByID(ctx, command.UserID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if err := h.userRepo.Delete(ctx, command.UserID); err != nil {
		return nil, err
	}

	return &CurrentUserDeleteCommandResponse{
		Message: "User deleted successfully",
	}, nil
}
