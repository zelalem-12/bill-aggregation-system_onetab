package middleware

import "github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/infrastructure/config"

type ProviderMiddleware struct {
	config *config.Config
	Middleware
}

func NewProviderMiddleware(config *config.Config) *ProviderMiddleware {
	return &ProviderMiddleware{
		config:     config,
		Middleware: Middleware{config: config},
	}
}
