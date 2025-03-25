package handler

import (
	"fmt"
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
)

type HomeHandler struct {
	config *config.Config
	tmpl   *template.Template
}

func NewHomeHandler(config *config.Config) (*HomeHandler, error) {

	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		return nil, err
	}

	return &HomeHandler{
		config: config,
		tmpl:   tmpl,
	}, nil
}

func (h *HomeHandler) Home(c echo.Context) error {

	data := map[string]interface{}{
		"userServiceDocsUrl":     fmt.Sprintf("%s:%d/api/v1/user/swagger/index.html", "http://127.0.0.1", 8000),
		"providerServiceDocsUrl": fmt.Sprintf("%s:%d/api/v1/provider/swagger/index.html", "http://127.0.0.1", 8000),
		"billServiceDocsUrl":     fmt.Sprintf("%s:%d/api/v1/bill/swagger/index.html", "http://127.0.0.1", 8000),
	}

	return h.tmpl.Execute(c.Response(), data)
}
