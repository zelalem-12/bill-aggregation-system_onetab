package passwordresetrequest

import "github.com/go-playground/validator/v10"

type PasswordResetRequestCommand struct {
	Email string `json:"email" validate:"required,email"`
}

func (passwordResetRequestCommand *PasswordResetRequestCommand) Validate() error {
	validate := validator.New()

	return validate.Struct(passwordResetRequestCommand)
}

type PasswordResetRequestCommandResponse struct {
	Message string `json:"message"`
}
