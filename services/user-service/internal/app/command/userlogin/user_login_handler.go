package userlogin

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/zelalem-12/onetab/internal/app/repo"
	"github.com/zelalem-12/onetab/internal/domain"
	"github.com/zelalem-12/onetab/internal/infrastructure/config"
	"github.com/zelalem-12/onetab/internal/util"
)

type UserLoginCommandHandler struct {
	config    *config.Config
	userRepo  repo.UserRepo
	tokenRepo repo.TokenRepo
}

func NewUserLoginCommandHandler(
	cfg *config.Config,
	UserRepo repo.UserRepo,
	tokenRepo repo.TokenRepo,
) *UserLoginCommandHandler {
	return &UserLoginCommandHandler{
		config:    cfg,
		userRepo:  UserRepo,
		tokenRepo: tokenRepo,
	}
}

func (h *UserLoginCommandHandler) Handle(ctx context.Context, command *UserLoginCommand) (*UserLoginCommandResponse, error) {

	user, err := h.userRepo.FindByEmail(ctx, command.Email)

	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(user.GetID())
	if err != nil {
		return nil, err
	}

	isMatched := util.ComparePasswords(command.Password, user.GetPassword())

	if !isMatched {
		return nil, errors.New("wrong password")
	}

	accessToken, err := util.GenerateAccessToken(h.config.ACCESS_TOKEN_KEY, user, h.config.ACCESS_TOKEN_EXPIRY)
	if err != nil {
		return nil, err
	}

	refreshToken, err := util.GenerateNonAccessToken(h.config.REFRESH_TOKEN_KEY, userID, h.config.REFRESH_TOKEN_EXPIRY)
	if err != nil {
		return nil, err
	}

	refreshTokenData := domain.NewToken(refreshToken, user.GetID())

	_, err = h.tokenRepo.Save(ctx, refreshTokenData)
	if err != nil {
		return nil, err
	}

	return &UserLoginCommandResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
