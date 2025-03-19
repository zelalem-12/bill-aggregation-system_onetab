package app

import (
	"github.com/mehdihadeli/go-mediatr"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/command/addprovider"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providerbyid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providerbyname"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providers"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/repo"

	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/infrastructure/config"
)

func RegisterCQRSHandlers(
	cfg *config.Config,
	providerRepo repo.ProviderRepo,
) error {

	addproviderCommandHandler := addprovider.NewAddProviderCommandHandler(providerRepo)

	providerByIdQueryHandler := providerbyid.NewGetProviderByIDQueryHandler(providerRepo)
	providerByNameQueryHandler := providerbyname.NewGetProviderByNameQueryHandler(providerRepo)
	providersQueryHandler := providers.NewGetProvidersQueryHandler(providerRepo)

	if err := mediatr.RegisterRequestHandler(addproviderCommandHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(providerByIdQueryHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(providerByNameQueryHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(providersQueryHandler); err != nil {
		return err
	}

	return nil
}
