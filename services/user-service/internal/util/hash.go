package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/zelalem-12/onetab/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateAccessToken(accessSecretKey string, user *domain.User, expiryTime string) (string, error) {

	userID, err := uuid.Parse(user.GetID())
	if err != nil {
		return "", err
	}

	duration, err := ParseExpiryTime(expiryTime)
	if err != nil {
		return "", err
	}

	claims := CustomClaims{
		User: BasicUserInfo{
			UserID:    userID,
			Email:     user.GetEmail(),
			FirstName: user.GetFirstName(),
			LastName:  user.GetLastName(),
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Issuer:    "onetab",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(accessSecretKey))
}

func GenerateNonAccessToken(secretKey string, userID uuid.UUID, expiryTime string) (string, error) {

	duration, err := ParseExpiryTime(expiryTime)
	if err != nil {
		return "", err
	}

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Issuer:    "onetab",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
}

func ParseAndValidateNonAccessToken(tokenString, secret string) (*uuid.UUID, error) {

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid || claims.UserID == uuid.Nil {
		return nil, err
	}

	return &claims.UserID, nil
}

func ParseExpiryTime(expiryTime string) (time.Duration, error) {

	unit := expiryTime[len(expiryTime)-1]

	if unit == 's' || unit == 'm' || unit == 'h' {
		return time.ParseDuration(expiryTime)
	}

	if unit == 'd' {
		value, err := strconv.Atoi(expiryTime[:len(expiryTime)-1])
		if err != nil {
			return 0, err
		}

		return time.ParseDuration(fmt.Sprintf("%dh", value*24))
	}

	if unit == 'w' {
		value, err := strconv.Atoi(expiryTime[:len(expiryTime)-1])
		if err != nil {
			return 0, err
		}

		return time.ParseDuration(fmt.Sprintf("%dh", value*24*7))
	}

	if unit == 'M' {
		value, err := strconv.Atoi(expiryTime[:len(expiryTime)-1])
		if err != nil {
			return 0, err
		}

		return time.ParseDuration(fmt.Sprintf("%dh", value*24*30))
	}

	return time.ParseDuration(expiryTime)
}
