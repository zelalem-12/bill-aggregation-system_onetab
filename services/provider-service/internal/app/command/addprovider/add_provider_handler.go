package addprovider

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/domain"
)

type AddProviderCommandHandler struct {
	ProviderRepo repo.ProviderRepo
}

func NewAddProviderCommandHandler(providerRepo repo.ProviderRepo) *AddProviderCommandHandler {
	return &AddProviderCommandHandler{ProviderRepo: providerRepo}
}

func (h *AddProviderCommandHandler) Handle(ctx context.Context, cmd *AddProviderCommand) (*AddProviderCommandResponse, error) {

	exixtingProvider, err := h.ProviderRepo.FindByName(ctx, cmd.Name)
	if err != nil {
		return nil, fmt.Errorf("error getting provider by name: %v", err)
	}

	if exixtingProvider != nil {
		providerId, err := uuid.Parse(exixtingProvider.GetID())
		if err != nil {
			providerId = uuid.New()
		}
		return &AddProviderCommandResponse{
			ProviderID: providerId,
		}, nil
	}

	provider := domain.NewProvider(cmd.Name, cmd.ClientID, cmd.ClientSecret, cmd.TokenURL, cmd.APIURL)

	persistedProvider, err := h.ProviderRepo.Save(ctx, provider)
	if err != nil {
		return nil, fmt.Errorf("error saving provider: %v", err)
	}

	providerID, err := uuid.Parse(persistedProvider.GetID())
	if err != nil {
		return nil, fmt.Errorf("error parsing linked account ID: %v", err)
	}

	return &AddProviderCommandResponse{
		ProviderID: providerID,
	}, nil
}
