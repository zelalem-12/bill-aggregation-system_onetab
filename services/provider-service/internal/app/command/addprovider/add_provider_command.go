package addprovider

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AddProviderCommand struct {
	Name         string `json:"name" validate:"required"`
	ClientID     string `json:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
	TokenURL     string `json:"token_url" validate:"required"`
	APIURL       string `json:"api_url" validate:"required"`
}

func (c *AddProviderCommand) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

type AddProviderCommandResponse struct {
	ProviderID uuid.UUID `json:"provider_id"`
}
