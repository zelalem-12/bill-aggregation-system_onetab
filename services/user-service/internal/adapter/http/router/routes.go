package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/zelalem-12/onetab/internal/adapter/http/handler"
	"github.com/zelalem-12/onetab/internal/adapter/http/middleware"
	"github.com/zelalem-12/onetab/internal/infrastructure/config"
)

func RegisterHomeRoute(
	e *echo.Echo,
	config *config.Config,
	homeHandler *handler.HomeHandler,
) {
	e.GET("", homeHandler.Home)
}
func RegisterSwaggerRoute(baseApi *echo.Group) {
	baseApi.GET("/swagger/*filepath", echoSwagger.WrapHandler)
}

func RegisterAuthRoutes(
	baseApi *echo.Group,
	authMiddleware *middleware.AuthMiddleware,
	authHandler *handler.AuthHandler,
) {
	authRoutes := baseApi.Group("/auth")

	authRoutes.POST("/register", authHandler.SignupUserHandler)
	authRoutes.POST("/verify-email", authHandler.VerifyEmailHandler, authMiddleware.ValidateVerifyTokenMiddleware)
	authRoutes.POST("/set-password", authHandler.ResetPasswordHandler, authMiddleware.ValidateVerifyTokenMiddleware)

	authRoutes.POST("/request-password-reset", authHandler.RequestPasswordResetHandler)
	authRoutes.POST("/reset-password", authHandler.ResetPasswordHandler, authMiddleware.ValidateVerifyTokenMiddleware)

	authRoutes.POST("/login", authHandler.LoginUserHandler)
	authRoutes.POST("/refresh-token", authHandler.RefreshTokenHandler, authMiddleware.ValidateRefreshTokenMiddleware)

}
