package addprovider

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AddProviderCommand struct {
	Name       string `json:"name" validate:"required"`
	APIURL     string `json:"api_url" validate:"required"`
	AuthMethod string `json:"auth_method" validate:"required"`

	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	TokenURL     string `json:"token_url"`

	APIToken string `json:"api_token"`
}

func (c *AddProviderCommand) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

type AddProviderCommandResponse struct {
	ProviderID uuid.UUID `json:"provider_id"`
}
