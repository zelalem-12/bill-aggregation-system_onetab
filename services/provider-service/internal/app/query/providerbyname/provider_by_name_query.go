package providerbyname

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProviderResponse struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	APIURL string    `json:"api_url"`
}

type GetProviderByNameQuery struct {
	ProviderName string `json:"provider_name" validate:"required"`
}

func (q *GetProviderByNameQuery) Validate() error {
	validate := validator.New()
	return validate.Struct(q)
}

type GetProviderByNameQueryResponse struct {
	Provider *ProviderResponse `json:"provider"`
}
