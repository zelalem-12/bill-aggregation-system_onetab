package users

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
)

type GetUsersQueryHandler struct {
	userRepo repo.UserRepo
}

func NewGetUsersQueryHandler(userRepo repo.UserRepo) *GetUsersQueryHandler {
	return &GetUsersQueryHandler{userRepo: userRepo}
}

func (h *GetUsersQueryHandler) Handle(ctx context.Context, query *UsersQuery) (*UsersQueryResponse, error) {

	usersData, err := h.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	if usersData == nil {
		return nil, errors.New("user not found")
	}

	users := make([]*User, 0)

	for _, userDomain := range usersData {

		user := &User{}

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

		accounts := make([]*LinkedAccount, 0)

		for _, linkedAccount := range userDomain.GetLinkedAccounts() {
			accounts = append(accounts, &LinkedAccount{
				ID:         linkedAccount.GetID(),
				ProviderID: linkedAccount.GetProviderID(),
				AuthToken:  linkedAccount.GetAuthToken(),
			})
		}

		user.LinkedAccounts = accounts

		users = append(users, user)
	}

	return &UsersQueryResponse{Users: users}, nil
}
