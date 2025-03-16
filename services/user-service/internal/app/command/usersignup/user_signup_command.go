package usersignup

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserSignupCommand struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

func (signupCommand *UserSignupCommand) Validate() error {
	validate := validator.New()

	return validate.Struct(signupCommand)
}

type UserSignupCommandResponse struct {
	UserID  uuid.UUID `json:"user_id"`
	Message string    `json:"message"`
}
