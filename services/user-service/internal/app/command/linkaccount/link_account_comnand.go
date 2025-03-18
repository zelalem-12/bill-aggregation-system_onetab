package linkaccount

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type LinkAccountCommand struct {
	UserID     uuid.UUID `validate:"required"`
	ProviderID uuid.UUID `validate:"required"`
	AuthToken  string    `validate:"required"`
}

func (c *LinkAccountCommand) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

type LinkAccountCommandResponse struct {
	AccountID uuid.UUID `json:"account_id"`
	Message   string    `json:"message"`
}
