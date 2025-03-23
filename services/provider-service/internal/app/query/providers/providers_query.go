package providers

import "github.com/google/uuid"

type ProviderResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	APIURL     string    `json:"api_url"`
	AUTHMethod string    `json:"auth_method"`
}

type GetProvidersQuery struct{}

type GetProvidersQueryResponse struct {
	Providers []*ProviderResponse `json:"providers"`
}
