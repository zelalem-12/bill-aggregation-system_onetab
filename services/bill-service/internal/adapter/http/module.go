package http

import (
	"github.com/labstack/echo/v4"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/infrastructure/config"
	"go.uber.org/fx"
)

func setupRoutes(
	config *config.Config,
	e *echo.Echo,

	//homeHandler *handler.HomeHandler,

) *echo.Group {

	//router.RegisterHomeRoute(e, config, homeHandler)

	v1 := e.Group("/api/v1")

	return v1
}

var Module = fx.Options(
	fx.Provide(
	//handler.NewHomeHandler,
	),
	fx.Invoke(setupRoutes),
)
