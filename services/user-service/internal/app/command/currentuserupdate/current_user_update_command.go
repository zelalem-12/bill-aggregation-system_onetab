package currentuserupdate

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CurrentUserUpdateCommand struct {
	UserID         uuid.UUID `json:"user_id" validate:"required"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	ProfilePicture string    `json:"profile_picture"`
}

func (signupCommand *CurrentUserUpdateCommand) Validate() error {
	validate := validator.New()
	return validate.Struct(signupCommand)
}

type CurrentUserUpdateCommandResponse struct {
	Message string `json:"message"`
}
