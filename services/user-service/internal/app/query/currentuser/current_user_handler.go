package currentuser

import (
	"context"
	"errors"

	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
)

type GetCurrentUserQueryHandler struct {
	userRepo repo.UserRepo
}

func NewCurrentUserQueryHandler(userRepo repo.UserRepo) *GetCurrentUserQueryHandler {
	return &GetCurrentUserQueryHandler{userRepo: userRepo}
}

func (h *GetCurrentUserQueryHandler) Handle(ctx context.Context, query *CurrentUserQuery) (*CurrentUserQueryResponse, error) {

	user, err := h.userRepo.FindByID(ctx, query.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return NewCurrentUserQueryResponse(user)
}
