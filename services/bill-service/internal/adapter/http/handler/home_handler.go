package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/infrastructure/config"
)

type HomeHandler struct {
	config *config.Config
}

func NewHomeHandler(config *config.Config) (*HomeHandler, error) {

	return &HomeHandler{
		config: config,
	}, nil
}

func (h *HomeHandler) Home(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, From Bill Service!")
}
