package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/linkaccount"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/unlinkaccount"
)

type LinkAccountRequest struct {
	ProviderID     uuid.UUID `json:"provider_id" validate:"required"`
	AuthToken      string    `json:"auth_token" validate:"required"`
	RefreshToken   string    `json:"refresh_token" validate:"required"`
	TokenType      string    `json:"token_type" validate:"required"`
	ProviderUserID string    `json:"provider_user_id" validate:"required"`
	ExpiresAt      time.Time `json:"expires_at" validate:"required"`
}

func (r *LinkAccountRequest) Validate() error {
	return validate.Struct(r)
}

func (r *LinkAccountRequest) ToCommand() *linkaccount.LinkAccountCommand {
	return &linkaccount.LinkAccountCommand{
		ProviderID:     r.ProviderID,
		AuthToken:      r.AuthToken,
		RefreshToken:   r.RefreshToken,
		TokenType:      r.TokenType,
		ProviderUserID: r.ProviderUserID,
		ExpiresAt:      r.ExpiresAt,
	}
}

type UnlinkAccountRequest struct {
	AccountID uuid.UUID `param:"account_id" validate:"required,uuid"`
}

func (r *UnlinkAccountRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UnlinkAccountRequest) ToCommand() (*unlinkaccount.UnlinkAccountCommand, error) {

	return &unlinkaccount.UnlinkAccountCommand{
		AccountID: r.AccountID,
	}, nil
}
