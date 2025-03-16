package router

import (
	"github.com/labstack/echo/v4"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/adapter/http/handler"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/infrastructure/config"
)

func RegisterHomeRoute(
	e *echo.Echo,
	config *config.Config,
	homeHandler *handler.HomeHandler,
) {
	e.GET("", homeHandler.Home)
}
