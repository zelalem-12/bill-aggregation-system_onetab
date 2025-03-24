package client

import (
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/domain"
)

type ProviderServiceClient interface {
	FetchBillsFromProvider(account *LinkedAccount, provider *domain.Provider) ([]*ProviderBillResponse, error)
}
