package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	binder   *echo.DefaultBinder
	validate *validator.Validate
}

func NewHandler() Handler {
	return Handler{
		binder:   &echo.DefaultBinder{},
		validate: validator.New(),
	}
}

func (h Handler) BindAndValidate(c echo.Context, i interface{}) error {
	if i == nil {
		return nil
	}

	if err := h.binder.BindHeaders(c, i); err != nil {
		return err
	}

	if err := h.binder.BindBody(c, i); err != nil {
		return err
	}

	if err := h.binder.BindPathParams(c, i); err != nil {
		return err
	}

	if err := h.binder.BindQueryParams(c, i); err != nil {
		return err
	}

	if err := h.validate.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
