package http

import (
	"github.com/labstack/echo/v4"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/http/handler"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/http/middleware"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/http/router"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
	"go.uber.org/fx"
)

func setupRoutes(
	config *config.Config,
	e *echo.Echo,

	authMiddleware *middleware.AuthMiddleware,

	homeHandler *handler.HomeHandler,

	authHandler *handler.AuthHandler,

) *echo.Group {

	router.RegisterHomeRoute(e, config, homeHandler)

	v1 := e.Group("/api/v1")

	router.RegisterSwaggerRoute(v1)
	router.RegisterAuthRoutes(v1, authMiddleware, authHandler)

	return v1
}

var Module = fx.Options(
	fx.Provide(
		middleware.NewAuthMiddleware,

		handler.NewHomeHandler,
		handler.NewAuthHandler,
	),
	fx.Invoke(setupRoutes),
)
