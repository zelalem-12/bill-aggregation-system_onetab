package providers

import (
	"context"

	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/service"
)

type GetProvidersQueryHandler struct {
	ProviderRepo repo.ProviderRepo
}

func NewGetProvidersQueryHandler(providerRepo repo.ProviderRepo) *GetProvidersQueryHandler {
	return &GetProvidersQueryHandler{
		ProviderRepo: providerRepo,
	}
}

func (h *GetProvidersQueryHandler) Handle(ctx context.Context, query *GetProvidersQuery) (*GetProvidersQueryResponse, error) {

	providers, err := h.ProviderRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	providerResponses := make([]*ProviderResponse, 0, len(providers))
	for _, provider := range providers {
		providerID, err := service.ToUUID(provider.GetID())
		if err != nil {
			return nil, err
		}

		providerResponse := ProviderResponse{
			ID:         providerID,
			Name:       provider.GetName(),
			APIURL:     provider.GetAPIURL(),
			AUTHMethod: provider.GetAuthMethod(),
		}

		providerResponses = append(providerResponses, &providerResponse)
	}

	return &GetProvidersQueryResponse{
		Providers: providerResponses,
	}, nil
}
