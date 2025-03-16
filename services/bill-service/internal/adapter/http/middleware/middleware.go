package middleware

import "github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/infrastructure/config"

type Middleware struct {
	config *config.Config
}

func NewMiddleware(config *config.Config) *Middleware {
	return &Middleware{
		config: config,
	}
}
