package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/domain"
)

type Token struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;not null;default:uuid_generate_v4()"`
	Token     string    `gorm:"not null;unique"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;unique"`
	User      *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ExpiresAt time.Time `gorm:"not null"`
}

func (t *Token) FromDomainModel(token *domain.Token) {

	if token.GetID() != "" {
		tokenID, err := uuid.Parse(token.GetID())
		if err != nil {
			return
		}
		t.ID = tokenID
	}

	t.Token = token.GetToken()

	userID, err := uuid.Parse(token.GetUserID())
	if err != nil {
		return
	}
	t.UserID = userID

}

func (t *Token) ToDomainModel() *domain.Token {
	token := domain.Token{}

	token.SetID(t.ID.String())
	token.SetToken(t.Token)
	token.SetUserID(t.UserID.String())

	return &token
}
