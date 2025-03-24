package client

import (
	"time"

	"github.com/google/uuid"
)

type LinkedAccount struct {
	ID             uuid.UUID `json:"id"`
	ProviderID     uuid.UUID `json:"provider_id"`
	AuthToken      string    `json:"auth_token"`
	RefreshToken   string    `json:"refresh_token"`
	ExpiresAt      time.Time `json:"expires_at"`
	TokenType      string    `json:"token_type"`
	ProviderUserID string    `json:"provider_user_id"`
}

type UserDetail struct {
	ID             uuid.UUID        `json:"id"`
	FirstName      string           `json:"first_name"`
	LastName       string           `json:"last_name"`
	Email          string           `json:"email"`
	IsVerified     bool             `json:"is_verified"`
	ProfilePicture string           `json:"profile_picture"`
	LinkedAccounts []*LinkedAccount `json:"linked_accounts"`
}

type UsersResponse struct {
	Users []*UserDetail `json:"users"`
}

type UserServiceClient interface {
	GetUsers() (*UsersResponse, error)
	GetUserDetail(userID uuid.UUID) (*UserDetail, error)
}
