package userlogin

import (
	"github.com/go-playground/validator/v10"
)

type UserLoginCommand struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (loginCommand *UserLoginCommand) Validate() error {
	validate := validator.New()

	return validate.Struct(loginCommand)
}

type UserLoginCommandResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
