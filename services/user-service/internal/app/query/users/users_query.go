package users

import (
	"github.com/google/uuid"
)

type UsersQuery struct{}

type LinkedAccount struct {
	ID         string `json:"id"`
	ProviderID string `json:"provider_id"`
	AuthToken  string `json:"auth_token"`
}

type User struct {
	ID             uuid.UUID        `json:"id"`
	FirstName      string           `json:"first_name"`
	LastName       string           `json:"last_name"`
	Email          string           `json:"email"`
	IsVerified     bool             `json:"is_verified"`
	ProfilePicture string           `json:"profile_picture"`
	LinkedAccounts []*LinkedAccount `json:"linked_accounts"`
}

type UsersQueryResponse struct {
	Users []*User `json:"users"`
}
