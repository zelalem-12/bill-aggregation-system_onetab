package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/util"
)

type AuthMiddleware struct {
	config *config.Config
}

func NewAuthMiddleware(config *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		config: config,
	}
}

func (m *AuthMiddleware) ValidateAccessTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		accessToken := c.QueryParam("token")
		if accessToken == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "access token is required"})
		}

		userId, err := util.ParseAndValidateNonAccessToken(accessToken, m.config.ACCESS_TOKEN_KEY)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid access access token"})
		}

		c.Set("user_id", *userId)
		c.Set("token", accessToken)

		return next(c)
	}
}

func (m *AuthMiddleware) ValidateRefreshTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		refreshToken := c.QueryParam("token")
		if refreshToken == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "refresh token is required"})
		}

		userId, err := util.ParseAndValidateNonAccessToken(refreshToken, m.config.REFRESH_TOKEN_KEY)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid refresh token"})
		}

		c.Set("user_id", *userId)
		c.Set("token", refreshToken)

		return next(c)
	}
}
