package providerbyid

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProviderResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	APIURL     string    `json:"api_url"`
	AUTHMethod string    `json:"auth_method"`
}

type GetProviderByIDQuery struct {
	ProviderID uuid.UUID `json:"provider_id" validate:"required"`
}

func (q *GetProviderByIDQuery) Validate() error {
	validate := validator.New()
	return validate.Struct(q)
}

type GetProviderByIDQueryResponse struct {
	Provider *ProviderResponse `json:"provider"`
}
