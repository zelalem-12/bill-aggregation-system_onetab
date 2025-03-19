package middleware

import "github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/infrastructure/config"

type BillMiddleware struct {
	config *config.Config
	Middleware
}

func NewBillMiddleware(config *config.Config) *BillMiddleware {
	return &BillMiddleware{
		config:     config,
		Middleware: Middleware{config: config},
	}
}
