package userlogout

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserLogoutCommand struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	Token  string    `json:"token" validate:"required"`
}

func (lc *UserLogoutCommand) Validate() error {
	validate := validator.New()

	return validate.Struct(lc)
}

type UserLogoutCommandResponse struct {
	Message string `json:"message"`
}
