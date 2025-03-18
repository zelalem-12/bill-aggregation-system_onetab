package currentuserupdate

import (
	"context"
	"errors"

	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
)

type UpdateCurrentUserHandler struct {
	userRepo repo.UserRepo
}

func NewUpdateCurrentUserHandler(userRepo repo.UserRepo) *UpdateCurrentUserHandler {
	return &UpdateCurrentUserHandler{userRepo: userRepo}
}

func (h *UpdateCurrentUserHandler) Handle(ctx context.Context, command *CurrentUserUpdateCommand) (*CurrentUserUpdateCommandResponse, error) {

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

	if command.FirstName != "" {
		user.SetFirstName(command.FirstName)
	}
	if command.LastName != "" {
		user.SetLastName(command.LastName)
	}

	if command.ProfilePicture != "" {
		user.SetProfilePicture(command.ProfilePicture)
	}

	if _, err := h.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return &CurrentUserUpdateCommandResponse{
		Message: "User updated successfully",
	}, nil
}
