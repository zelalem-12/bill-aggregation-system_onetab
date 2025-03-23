package currentuser

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/domain"
)

type CurrentUserQuery struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type LinkedAccount struct {
	ID         string `json:"id"`
	ProviderID string `json:"provider_id"`
	AuthToken  string `json:"auth_token"`
}

func (q *CurrentUserQuery) Validate() error {
	validate := validator.New()
	return validate.Struct(q)
}

type CurrentUserQueryResponse struct {
	ID             uuid.UUID       `json:"id"`
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	Email          string          `json:"email"`
	IsVerified     bool            `json:"is_verified"`
	ProfilePicture string          `json:"profile_picture"`
	LinkedAccounts []LinkedAccount `json:"linked_accounts"`
}

func NewCurrentUserQueryResponse(userDomain *domain.User) (*CurrentUserQueryResponse, error) {
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
			ID:         linkedAccount.GetID(),
			ProviderID: linkedAccount.GetProviderID(),
			AuthToken:  linkedAccount.GetAuthToken(),
		})
	}

	user.LinkedAccounts = accounts

	return &user, nil

}
