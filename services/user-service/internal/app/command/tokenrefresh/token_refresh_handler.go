package tokenrefresh

import (
	"context"

	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/util"
)

type TokenRefreshCommandHandler struct {
	config    *config.Config
	userRepo  repo.UserRepo
	tokenRepo repo.TokenRepo
}

func NewTokenRefreshCommandHandler(
	cfg *config.Config,
	userRepo repo.UserRepo,
	tokenRepo repo.TokenRepo,
) *TokenRefreshCommandHandler {
	return &TokenRefreshCommandHandler{
		config:    cfg,
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (h *TokenRefreshCommandHandler) Handle(ctx context.Context, command *TokenRefreshCommand) (*TokenRefreshCommandResponse, error) {

	tokenData, err := h.tokenRepo.Find(ctx, command.UserID, command.RefreshToken)
	if err != nil || tokenData == nil {
		return nil, err
	}

	user, err := h.userRepo.FindByID(ctx, command.UserID)
	if err != nil {
		return nil, err
	}

	accessToken, err := util.GenerateAccessToken(h.config.ACCESS_TOKEN_KEY, user, h.config.ACCESS_TOKEN_EXPIRY)
	if err != nil {
		return nil, err
	}

	return &TokenRefreshCommandResponse{
		AccessToken: accessToken,
	}, nil
}
