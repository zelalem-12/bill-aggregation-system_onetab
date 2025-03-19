package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providerbyid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providerbyname"
)

type GetProviderByIDRequest struct {
	ProviderID uuid.UUID `param:"provider_id" validate:"required"`
}

func (r *GetProviderByIDRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetProviderByIDRequest) ToQuery() *providerbyid.GetProviderByIDQuery {
	return &providerbyid.GetProviderByIDQuery{
		ProviderID: r.ProviderID,
	}
}

type GetProviderByNameRequest struct {
	ProviderName string `param:"provider_name" validate:"required"`
}

func (r *GetProviderByNameRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetProviderByNameRequest) ToQuery() *providerbyname.GetProviderByNameQuery {
	return &providerbyname.GetProviderByNameQuery{
		ProviderName: r.ProviderName,
	}
}
