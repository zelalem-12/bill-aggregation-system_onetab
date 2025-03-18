package unlinkaccount

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UnlinkAccountCommand struct {
	AccountID uuid.UUID `json:"account_id" validate:"required"`
}

func (c *UnlinkAccountCommand) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

type UnlinkAccountCommandResponse struct {
	AccountID uuid.UUID `json:"account_id"`
	Message   string    `json:"message"`
}
