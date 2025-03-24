package app

import (
	"github.com/mehdihadeli/go-mediatr"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/command/addprovider"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/command/refreshbills"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/command/syncbills"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providerbyid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providerbyname"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providers"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/repo"

	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/infrastructure/config"
)

func RegisterCQRSHandlers(
	cfg *config.Config,
	providerRepo repo.ProviderRepo,
	userClient client.UserServiceClient,
	providerClient client.ProviderServiceClient,
	billClient client.BillServiceClient,
) error {

	addproviderCommandHandler := addprovider.NewAddProviderCommandHandler(providerRepo)
	refreshbillsCommandHandler := refreshbills.NewRefreshBillsCommandHandler(providerRepo, userClient, providerClient, billClient)
	syncbillsCommandHandler := syncbills.NewSyncAllBillsCommandHandler(providerRepo, userClient, providerClient, billClient)

	providerByIdQueryHandler := providerbyid.NewGetProviderByIDQueryHandler(providerRepo)
	providerByNameQueryHandler := providerbyname.NewGetProviderByNameQueryHandler(providerRepo)
	providersQueryHandler := providers.NewGetProvidersQueryHandler(providerRepo)

	if err := mediatr.RegisterRequestHandler(addproviderCommandHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(refreshbillsCommandHandler); err != nil {
		return err
	}

	if err := mediatr.RegisterRequestHandler(syncbillsCommandHandler); err != nil {
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
