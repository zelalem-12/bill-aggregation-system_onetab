package util

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type BasicUserInfo struct {
	UserID    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type CustomClaims struct {
	User BasicUserInfo `json:"user"`
	jwt.RegisteredClaims
}

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}
