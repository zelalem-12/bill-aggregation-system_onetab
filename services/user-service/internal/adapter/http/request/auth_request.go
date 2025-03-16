package request

import (
	"github.com/zelalem-12/onetab/internal/app/command/passwordreset"
	"github.com/zelalem-12/onetab/internal/app/command/passwordresetrequest"
	"github.com/zelalem-12/onetab/internal/app/command/userlogin"
	"github.com/zelalem-12/onetab/internal/app/command/usersignup"
	"github.com/zelalem-12/onetab/internal/util"
)

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

func (r *RegisterRequest) Validate() error {

	return validate.Struct(r)
}

func (r *RegisterRequest) ToCommand() *usersignup.UserSignupCommand {
	return &usersignup.UserSignupCommand{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
	}
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (l *LoginRequest) Validate() error {

	return validate.Struct(l)
}
func (l *LoginRequest) ToCommand() *userlogin.UserLoginCommand {
	return &userlogin.UserLoginCommand{
		Email:    l.Email,
		Password: l.Password,
	}
}

type PasswordResetRequestRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (r *PasswordResetRequestRequest) Validate() error {

	return validate.Struct(r)
}
func (r *PasswordResetRequestRequest) ToCommand() *passwordresetrequest.PasswordResetRequestCommand {
	return &passwordresetrequest.PasswordResetRequestCommand{
		Email: r.Email,
	}
}

type PasswordResetRequest struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (r *PasswordResetRequest) Validate() error {

	if r.Password != r.ConfirmPassword {
		return util.CreateError("password and confirm password do not match")
	}
	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}
func (r *PasswordResetRequest) ToCommand() *passwordreset.PasswordResetCommand {
	return &passwordreset.PasswordResetCommand{
		Password: r.Password,
	}
}
