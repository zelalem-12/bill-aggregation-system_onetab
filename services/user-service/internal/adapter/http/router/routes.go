package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/http/handler"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/http/middleware"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
)

func RegisterHomeRoute(
	e *echo.Echo,
	config *config.Config,
	homeHandler *handler.HomeHandler,
) {
	e.GET("", homeHandler.Home)
}

func RegisterinternalUserRoutes(
	baseApi *echo.Echo,
	userHandler *handler.UserHandler) {

	internalRoute := baseApi.Group("/internal")
	internalRoute.GET("/users", userHandler.GetUsersHandler)
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
	authRoutes.POST("/verify-email", authHandler.VerifyEmailHandler, authMiddleware.ValidateAccessTokenMiddleware)
	authRoutes.POST("/set-password", authHandler.SetPasswordHandler, authMiddleware.ValidateAccessTokenMiddleware)

	authRoutes.POST("/request-password-reset", authHandler.RequestPasswordResetHandler)
	authRoutes.POST("/reset-password", authHandler.ResetPasswordHandler, authMiddleware.ValidateAccessTokenMiddleware)

	authRoutes.POST("/login", authHandler.LoginUserHandler)
	authRoutes.POST("/refresh-token", authHandler.RefreshTokenHandler, authMiddleware.ValidateRefreshTokenMiddleware)

	authRoutes.POST("/logout", authHandler.LogoutHandler, authMiddleware.ValidateAccessTokenMiddleware)
}

func RegisterUserRoutes(
	baseApi *echo.Group,
	userMiddleware *middleware.UserMiddleware,
	userHandler *handler.UserHandler,
) {
	userRoutes := baseApi.Group("/user", userMiddleware.ConstructJWTConfig(), userMiddleware.AttachCustomClaims)

	userRoutes.POST("/change-password", userHandler.ChangePasswordHandler)
	userRoutes.GET("/me", userHandler.GetCurrentUserHandler)
	userRoutes.PUT("/me", userHandler.UpdateCurrentUserHandler)
	userRoutes.DELETE("/me", userHandler.DeleteCurrentUserHandler)
}

func RegisterLinkedAccountRoutes(
	baseApi *echo.Group,
	linkedAccountMiddleware *middleware.LinkedAccountMiddleware,
	linkedAccountHandler *handler.LinkedAccountHandler,
) {
	linkedAccountRoutes := baseApi.Group("/accounts", linkedAccountMiddleware.ConstructJWTConfig(), linkedAccountMiddleware.AttachCustomClaims)

	linkedAccountRoutes.POST("/link", linkedAccountHandler.LinkAccountHandler)
	linkedAccountRoutes.DELETE("/:account_id", linkedAccountHandler.UnlinkAccountHandler)
}
