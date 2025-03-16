package passwordreset

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PasswordResetCommand struct {
	UserID     uuid.UUID `json:"user_id" validate:"required"`
	Password   string    `json:"password" validate:"required"`
	ResetToken string    `json:"reset_token" validate:"required"`
}

func (passwordResetCommand *PasswordResetCommand) Validate() error {
	validate := validator.New()
	return validate.Struct(passwordResetCommand)
}

type PasswordResetCommandResponse struct {
	Message string `json:"message"`
}
