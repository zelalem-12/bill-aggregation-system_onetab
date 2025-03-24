package currentuser

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
)

type GetCurrentUserQueryHandler struct {
	userRepo repo.UserRepo
}

func NewCurrentUserQueryHandler(userRepo repo.UserRepo) *GetCurrentUserQueryHandler {
	return &GetCurrentUserQueryHandler{userRepo: userRepo}
}

func (h *GetCurrentUserQueryHandler) Handle(ctx context.Context, query *CurrentUserQuery) (*CurrentUserQueryResponse, error) {

	userDomain, err := h.userRepo.FindByID(ctx, query.UserID)
	if err != nil {
		return nil, err
	}
	if userDomain == nil {
		return nil, errors.New("user not found")
	}

	user := CurrentUserQueryResponse{}

	userID, err := uuid.Parse(userDomain.GetID())
	if err != nil {
		return nil, err
	}

	user.ID = userID
	user.FirstName = userDomain.GetFirstName()
	user.LastName = userDomain.GetLastName()
	user.Email = userDomain.GetEmail()
	user.IsVerified = userDomain.GetIsVerified()
	user.ProfilePicture = userDomain.GetProfilePicture()

	accounts := make([]LinkedAccount, 0)

	for _, linkedAccount := range userDomain.GetLinkedAccounts() {
		accounts = append(accounts, LinkedAccount{
			ID:             linkedAccount.GetID(),
			ProviderID:     linkedAccount.GetProviderID(),
			AuthToken:      linkedAccount.GetAuthToken(),
			RefreshToken:   linkedAccount.GetRefreshToken(),
			ExpiresAt:      linkedAccount.GetExpiresAt(),
			TokenType:      linkedAccount.GetTokenType(),
			ProviderUserID: linkedAccount.GetProviderUserID(),
		})
	}

	user.LinkedAccounts = accounts

	return &user, nil
}
