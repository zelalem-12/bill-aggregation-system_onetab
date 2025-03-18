package middleware

import "github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"

type UserMiddleware struct {
	config *config.Config
	Middleware
}

func NewUserMiddleware(config *config.Config) *UserMiddleware {
	return &UserMiddleware{
		config:     config,
		Middleware: Middleware{config: config},
	}
}
