package domain

import "time"

type LinkedAccount struct {
	Base
	userID         string
	user           *User
	providerID     string
	authToken      string
	refreshToken   string
	expiresAt      time.Time
	tokenType      string
	providerUserID string
}

func (la *LinkedAccount) SetUserID(userID string) {
	la.userID = userID
}

func (la *LinkedAccount) GetUserID() string {
	return la.userID
}

func (la *LinkedAccount) SetProviderID(providerID string) {
	la.providerID = providerID
}

func (la *LinkedAccount) GetProviderID() string {
	return la.providerID
}

func (account *LinkedAccount) SetAuthToken(authToken string) {
	account.authToken = authToken
}

func (la *LinkedAccount) GetAuthToken() string {
	return la.authToken
}

func (la *LinkedAccount) SetUser(user *User) {
	la.user = user
}

func (la *LinkedAccount) GetUser() *User {
	return la.user
}

func (la *LinkedAccount) SetRefreshToken(refreshToken string) {
	la.refreshToken = refreshToken
}

func (la *LinkedAccount) GetRefreshToken() string {
	return la.refreshToken
}

func (la *LinkedAccount) SetExpiresAt(expiresAt time.Time) {
	la.expiresAt = expiresAt
}

func (la *LinkedAccount) GetExpiresAt() time.Time {
	return la.expiresAt
}

func (la *LinkedAccount) SetTokenType(tokenType string) {
	la.tokenType = tokenType
}

func (la *LinkedAccount) GetTokenType() string {
	return la.tokenType
}

func (la *LinkedAccount) SetProviderUserID(providerUserID string) {
	la.providerUserID = providerUserID
}

func (la *LinkedAccount) GetProviderUserID() string {
	return la.providerUserID
}

func NewLinkedAccount(userID, providerID, authToken, refreshToken, tokenType, providerUserID string, expiresAt time.Time) *LinkedAccount {
	return &LinkedAccount{
		userID:         userID,
		providerID:     providerID,
		authToken:      authToken,
		refreshToken:   refreshToken,
		tokenType:      tokenType,
		providerUserID: providerUserID,
		expiresAt:      expiresAt,
	}
}
