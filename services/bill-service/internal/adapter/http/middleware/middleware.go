package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/infrastructure/config"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/util"
)

type Middleware struct {
	config *config.Config
}

func NewMiddleware(config *config.Config) *Middleware {
	return &Middleware{
		config: config,
	}
}

func (m *Middleware) ConstructJWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(m.config.ACCESS_TOKEN_KEY),
		NewClaimsFunc: func(c echo.Context) jwt.Claims { return &util.CustomClaims{} },
	})
}

func (m *Middleware) AttachCustomClaims(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token, ok := c.Get("user").(*jwt.Token)
		if !ok || token == nil {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*util.CustomClaims)
		if !ok || !token.Valid {
			return echo.ErrUnauthorized
		}

		c.Set("user", claims.User)

		return next(c)
	}
}
