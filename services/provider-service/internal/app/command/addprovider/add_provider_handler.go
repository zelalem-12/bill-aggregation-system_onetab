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

	if err := cmd.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}

	provider := domain.NewProvider(cmd.Name, cmd.ClientID, cmd.ClientSecret, cmd.TokenURL, cmd.APIURL)

	provider, err := h.ProviderRepo.Save(ctx, provider)
	if err != nil {
		return nil, fmt.Errorf("error saving provider: %v", err)
	}

	providerID, err := uuid.Parse(provider.GetID())
	if err != nil {
		return nil, fmt.Errorf("error parsing linked account ID: %v", err)
	}

	return &AddProviderCommandResponse{
		ProviderID: providerID,
	}, nil
}
