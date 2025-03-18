package currentuserdelete

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CurrentUserDeleteCommand struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

func (signupCommand *CurrentUserDeleteCommand) Validate() error {
	validate := validator.New()
	return validate.Struct(signupCommand)
}

type CurrentUserDeleteCommandResponse struct {
	Message string `json:"message"`
}
