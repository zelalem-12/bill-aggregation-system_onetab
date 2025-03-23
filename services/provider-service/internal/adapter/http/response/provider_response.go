package response

import (
	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providerbyid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providerbyname"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providers"
)

type ProviderResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	APIURL     string    `json:"api_url"`
	AUTHMethod string    `json:"auth_method"`
}

type GetProviderByIDResponse struct {
	Provider *ProviderResponse `json:"provider"`
}

func (r *GetProviderByIDResponse) FromQueryResponse(queryResponse *providerbyid.GetProviderByIDQueryResponse) {
	r.Provider = &ProviderResponse{
		ID:         queryResponse.Provider.ID,
		Name:       queryResponse.Provider.Name,
		APIURL:     queryResponse.Provider.APIURL,
		AUTHMethod: queryResponse.Provider.AUTHMethod,
	}
}

type GetProviderByNameResponse struct {
	Provider *ProviderResponse `json:"provider"`
}

func (r *GetProviderByNameResponse) FromQueryResponse(queryResponse *providerbyname.GetProviderByNameQueryResponse) {
	r.Provider = &ProviderResponse{
		ID:         queryResponse.Provider.ID,
		Name:       queryResponse.Provider.Name,
		APIURL:     queryResponse.Provider.APIURL,
		AUTHMethod: queryResponse.Provider.AUTHMethod,
	}
}

type GetProvidersResponse struct {
	Providers []*ProviderResponse `json:"providers"`
}

func (r *GetProvidersResponse) FromQueryResponse(queryResponse *providers.GetProvidersQueryResponse) {
	r.Providers = make([]*ProviderResponse, len(queryResponse.Providers))
	for i, provider := range queryResponse.Providers {
		r.Providers[i] = &ProviderResponse{
			ID:         provider.ID,
			Name:       provider.Name,
			APIURL:     provider.APIURL,
			AUTHMethod: provider.AUTHMethod,
		}
	}
}
