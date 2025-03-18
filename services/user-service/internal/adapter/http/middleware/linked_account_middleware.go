package middleware

import "github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"

type LinkedAccountMiddleware struct {
	config *config.Config
	Middleware
}

func NewLinkedAccountMiddleware(config *config.Config) *LinkedAccountMiddleware {
	return &LinkedAccountMiddleware{
		config:     config,
		Middleware: Middleware{config: config},
	}
}
