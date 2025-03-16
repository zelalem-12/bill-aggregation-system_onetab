package tokenrefresh

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TokenRefreshCommand struct {
	UserID       uuid.UUID `json:"user_id" validate:"required"`
	RefreshToken string    `json:"refresh_token" validate:"required"`
}

func (tokenRefreshCommand *TokenRefreshCommand) Validate() error {
	validate := validator.New()

	return validate.Struct(tokenRefreshCommand)
}

type TokenRefreshCommandResponse struct {
	AccessToken string `json:"access_token"`
}
