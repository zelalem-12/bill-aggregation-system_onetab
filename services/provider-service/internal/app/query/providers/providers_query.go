package providers

import "github.com/google/uuid"

type ProviderResponse struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	APIURL string    `json:"api_url"`
}

type GetProvidersQuery struct{}

type GetProvidersQueryResponse struct {
	Providers []*ProviderResponse `json:"providers"`
}
