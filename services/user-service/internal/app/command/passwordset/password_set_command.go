package passwordset

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PasswordSetCommand struct {
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	Password string    `json:"password" validate:"required"`
	SetToken string    `json:"set_token" validate:"required"`
}

func (ps *PasswordSetCommand) Validate() error {
	validate := validator.New()
	return validate.Struct(ps)
}

type PasswordSetCommandResponse struct {
	Message string `json:"message"`
}
