package http

import (
	"github.com/labstack/echo/v4"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/adapter/http/handler"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/adapter/http/middleware"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/adapter/http/router"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/infrastructure/config"
	"go.uber.org/fx"
)

func setupRoutes(
	config *config.Config,
	e *echo.Echo,

	billMiddleware *middleware.BillMiddleware,
	billHandler *handler.BillHandler,

) {
	v1 := e.Group("/api/v1")

	router.RegisterBillRoutes(v1, billMiddleware, billHandler)
	router.RegisterinternalBillRoutes(v1, billMiddleware, billHandler)

}

var Module = fx.Options(
	fx.Provide(
		middleware.NewBillMiddleware,
		handler.NewBillHandler,
	),
	fx.Invoke(setupRoutes),
)
