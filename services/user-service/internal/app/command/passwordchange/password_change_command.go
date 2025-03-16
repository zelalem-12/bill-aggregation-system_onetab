package passwordchange

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PasswordChangeCommand struct {
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	OldPassword string    `json:"old_password" validate:"required"`
	NewPassword string    `json:"new_password" validate:"required"`
}

func (passwordChangeCommand *PasswordChangeCommand) Validate() error {
	validate := validator.New()
	return validate.Struct(passwordChangeCommand)
}

type PasswordChangeCommandResponse struct {
	Message string `json:"message"`
}
