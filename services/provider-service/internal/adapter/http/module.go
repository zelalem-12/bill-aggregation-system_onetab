package http

import (
	"github.com/labstack/echo/v4"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/adapter/http/handler"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/adapter/http/middleware"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/adapter/http/router"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/infrastructure/config"
	"go.uber.org/fx"
)

func setupRoutes(
	config *config.Config,
	e *echo.Echo,

	providerMiddleware *middleware.ProviderMiddleware,
	providerHandler *handler.ProviderHandler,

) {

	v1 := e.Group("/api/v1")

	router.RegisterSwaggerRoute(v1)
	router.RegisterProviderRoutes(v1, providerMiddleware, providerHandler)

}

var Module = fx.Options(
	fx.Provide(
		middleware.NewProviderMiddleware,
		handler.NewProviderHandler,
	),
	fx.Invoke(setupRoutes),
)
