package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func defaultHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	c.Logger().Errorf("request error: %v", err)

	code := http.StatusInternalServerError
	message := any(http.StatusText(http.StatusInternalServerError))

	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code
		message = he.Message
	}

	var be *echo.BindingError
	if errors.As(err, &be) {
		code = be.Code
		message = be.Message
	}

	if m, ok := message.(string); ok {
		message = map[string]any{"message": m}
	}

	// Send response
	if c.Request().Method == http.MethodHead { // Issue #608
		err = c.NoContent(code)
	} else {
		err = c.JSON(code, message)
	}

	if err != nil {
		c.Logger().Errorf("failed writing error response: %v", err)
	}
}

func NewEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.Recover())

	e.HTTPErrorHandler = defaultHTTPErrorHandler

	return e
}

func ManageServerLifecycle(lc fx.Lifecycle, cfg *config.Config, db *gorm.DB, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			sqlDB, err := db.DB()
			if err != nil {
				return fmt.Errorf("failed to get raw database instance: %w", err)
			}
			if err := sqlDB.PingContext(ctx); err != nil {
				return fmt.Errorf("failed to connect to the database: %w", err)
			}

			go func() {
				if err := e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)); err != nil && err != http.ErrServerClosed {
					e.Logger.Fatal("shutting down the server")
				}
			}()
			return nil
		},

		OnStop: func(ctx context.Context) error {

			fmt.Println("Shutting down Echo server...")

			shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			err := e.Shutdown(shutdownCtx)
			if err != nil {
				return fmt.Errorf("failed to shutdown the server: %w", err)
			}

			sqlDB, err := db.DB()
			if err != nil {
				return fmt.Errorf("failed to get raw database instance for closing: %w", err)
			}

			if err := sqlDB.Close(); err != nil {
				return fmt.Errorf("failed to close database connection: %w", err)
			}

			return nil
		},
	})
}
