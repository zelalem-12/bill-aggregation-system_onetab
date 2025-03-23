package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/adapter/http/handler"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/adapter/http/middleware"
)

func RegisterinternalBillRoutes(
	baseApi *echo.Echo,
	billHandler *handler.BillHandler,
) {

	internalRoute := baseApi.Group("/internal")
	internalRoute.DELETE("/bills/:provider_id", billHandler.DeleteBillsByProvider)
}

func RegisterSwaggerRoute(baseApi *echo.Group) {
	baseApi.GET("/swagger/*filepath", echoSwagger.WrapHandler)
}

func RegisterBillRoutes(
	baseApi *echo.Group,
	billMiddleware *middleware.BillMiddleware,
	billHandler *handler.BillHandler) {

	billRoute := baseApi.Group("/bills", billMiddleware.ConstructJWTConfig(), billMiddleware.AttachCustomClaims)

	billRoute.GET("", billHandler.GetAggregatedBillsHandler)
	billRoute.GET("/:provider_name", billHandler.GetBillsByProviderHandler)
	billRoute.GET("/provider/:provider_id", billHandler.GetBillsByProviderIdHandler)
	billRoute.PUT("/:bill_id/pay", billHandler.MarkBillAsPaidHandler)
	billRoute.DELETE("/:bill_id", billHandler.DeleteBillHandler)
	billRoute.POST("", billHandler.CreateBillHandler)

	billRoute.GET("/overdue", billHandler.GetOverdueBillsHandler)
	billRoute.GET("/categories", billHandler.GetCategorySpendingHandler)
	billRoute.GET("/history", billHandler.GetBillPaymentHistoryHandler)
	billRoute.GET("/summary", billHandler.HandlerGetBillSummary)
	billRoute.GET("/summary/trends", billHandler.HandlerGetMonthlySpendingTrends)

}
