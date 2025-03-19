package router

import (
	"github.com/labstack/echo/v4"

	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/adapter/http/handler"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/adapter/http/middleware"
)

func RegisterProviderRoutes(
	baseApi *echo.Group,
	providerMiddleware *middleware.ProviderMiddleware,
	providerHandler *handler.ProviderHandler,
) {
	providerRoute := baseApi.Group("/providers", providerMiddleware.ConstructJWTConfig(), providerMiddleware.AttachCustomClaims)
	providerRoute.GET("/:provider_id", providerHandler.GetProviderByIdHandler)
	providerRoute.GET("/name/:provider_name", providerHandler.GetProviderByNameHandler)
	providerRoute.GET("", providerHandler.GetProvidersHandler)
}
