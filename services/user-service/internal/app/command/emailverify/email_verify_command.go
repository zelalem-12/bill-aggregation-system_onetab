package emailverify

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type EmailVerifyCommand struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	Token  string    `json:"token" validate:"required"`
}

func (emailVerifyCommand *EmailVerifyCommand) Validate() error {
	validate := validator.New()

	return validate.Struct(emailVerifyCommand)
}

type EmailVerifyCommandResponse struct {
	Message string `json:"message"`
}
