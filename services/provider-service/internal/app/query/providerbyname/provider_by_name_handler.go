package providerbyname

import (
	"context"
	"errors"

	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/service"
)

type GetProviderByNameQueryHandler struct {
	ProviderRepo repo.ProviderRepo
}

func NewGetProviderByNameQueryHandler(providerRepo repo.ProviderRepo) *GetProviderByNameQueryHandler {
	return &GetProviderByNameQueryHandler{
		ProviderRepo: providerRepo,
	}
}

func (h *GetProviderByNameQueryHandler) Handle(ctx context.Context, query *GetProviderByNameQuery) (*GetProviderByNameQueryResponse, error) {

	provider, err := h.ProviderRepo.FindByName(ctx, query.ProviderName)
	if err != nil {
		return nil, err
	}
	if provider == nil {
		return nil, errors.New("provider not found")
	}

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

	return &GetProviderByNameQueryResponse{
		Provider: &providerResponse,
	}, nil
}
