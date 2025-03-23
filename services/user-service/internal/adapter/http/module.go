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
	userMiddleware *middleware.UserMiddleware,
	linkedAccountMiddleware *middleware.LinkedAccountMiddleware,

	homeHandler *handler.HomeHandler,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	linkedAccountHandler *handler.LinkedAccountHandler,

) {
	router.RegisterHomeRoute(e, config, homeHandler)

	router.RegisterinternalUserRoutes(e, userHandler)

	v1 := e.Group("/api/v1")

	router.RegisterSwaggerRoute(v1)
	router.RegisterAuthRoutes(v1, authMiddleware, authHandler)
	router.RegisterUserRoutes(v1, userMiddleware, userHandler)
	router.RegisterLinkedAccountRoutes(v1, linkedAccountMiddleware, linkedAccountHandler)

}

var Module = fx.Options(
	fx.Provide(
		middleware.NewAuthMiddleware,
		middleware.NewUserMiddleware,
		middleware.NewLinkedAccountMiddleware,

		handler.NewHomeHandler,
		handler.NewAuthHandler,
		handler.NewUserHandler,
		handler.NewLinkedAccountHandler,
	),
	fx.Invoke(setupRoutes),
)
