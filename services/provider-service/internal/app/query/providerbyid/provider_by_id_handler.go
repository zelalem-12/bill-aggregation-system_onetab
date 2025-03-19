package providerbyid

import (
	"context"
	"errors"

	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/service"
)

type GetProviderByIDQueryHandler struct {
	ProviderRepo repo.ProviderRepo
}

func NewGetProviderByIDQueryHandler(providerRepo repo.ProviderRepo) *GetProviderByIDQueryHandler {
	return &GetProviderByIDQueryHandler{
		ProviderRepo: providerRepo,
	}
}

func (h *GetProviderByIDQueryHandler) Handle(ctx context.Context, query *GetProviderByIDQuery) (*GetProviderByIDQueryResponse, error) {
	provider, err := h.ProviderRepo.FindByID(ctx, query.ProviderID)
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
		ID:     providerID,
		Name:   provider.Name,
		APIURL: provider.APIURL,
	}

	return &GetProviderByIDQueryResponse{
		Provider: &providerResponse,
	}, nil
}
