package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/domain"
)

// refreshToken   string
// expiresAt      time.Time
// tokenType      string
// providerUserID string

type LinkedAccount struct {
	Base
	UserID         uuid.UUID `gorm:"not null"`
	User           *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ProviderID     uuid.UUID `gorm:"not null"`
	AuthToken      string    `gorm:"not null"`
	RefreshToken   string
	ExpiresAt      time.Time
	TokenType      string `gorm:"not null;default:Bearer"`
	ProviderUserID string
}

func (la *LinkedAccount) FromDomainModel(account *domain.LinkedAccount) error {

	if account.GetID() != "" {
		userID, err := uuid.Parse(account.GetID())
		if err != nil {
			return err
		}
		la.ID = userID
	}

	if account.GetUserID() != "" {
		userID, err := uuid.Parse(account.GetUserID())
		if err != nil {
			return err
		}
		la.UserID = userID
	}

	if account.GetProviderID() != "" {
		providerID, err := uuid.Parse(account.GetProviderID())
		if err != nil {
			return err
		}
		la.ProviderID = providerID
	}

	la.AuthToken = account.GetAuthToken()

	la.RefreshToken = account.GetRefreshToken()
	la.ExpiresAt = account.GetExpiresAt()
	la.TokenType = account.GetTokenType()
	la.ProviderUserID = account.GetProviderUserID()

	return nil

}

func (lc *LinkedAccount) ToDomainModel() *domain.LinkedAccount {
	domainAccount := domain.LinkedAccount{}

	domainAccount.SetID(lc.ID.String())
	domainAccount.SetUserID(lc.UserID.String())
	domainAccount.SetProviderID(lc.ProviderID.String())
	domainAccount.SetAuthToken(lc.AuthToken)
	domainAccount.SetRefreshToken(lc.RefreshToken)
	domainAccount.SetExpiresAt(lc.ExpiresAt)
	domainAccount.SetTokenType(lc.TokenType)
	domainAccount.SetProviderUserID(lc.ProviderUserID)

	domainAccount.SetCreatedAt(lc.CreatedAt)
	domainAccount.SetUpdatedAt(lc.UpdatedAt)
	if lc.DeletedAt.Valid {
		domainAccount.SetDeletedAt(&lc.DeletedAt.Time)
	}
	return &domainAccount

}
