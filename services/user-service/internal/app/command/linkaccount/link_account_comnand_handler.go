package linkaccount

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/domain"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
)

type LinkAccountCommandHandler struct {
	config            *config.Config
	linkedAccountRepo repo.LinkedAccountRepo
}

func NewLinkAccountCommandHandler(
	config *config.Config,
	linkedAccountRepo repo.LinkedAccountRepo,
) *LinkAccountCommandHandler {
	return &LinkAccountCommandHandler{
		config:            config,
		linkedAccountRepo: linkedAccountRepo,
	}
}

func (h *LinkAccountCommandHandler) Handle(ctx context.Context, command *LinkAccountCommand) (*LinkAccountCommandResponse, error) {
	if err := command.Validate(); err != nil {
		return nil, err
	}

	userID := command.UserID.String()
	providerID := command.ProviderID.String()

	linkedAccount := domain.NewLinkedAccount(
		userID,
		providerID,
		command.AuthToken,
		command.RefreshToken,
		command.TokenType,
		command.ProviderUserID,
		command.ExpiresAt,
	)

	savedLinkedAccount, err := h.linkedAccountRepo.Save(ctx, linkedAccount)
	if err != nil {
		return nil, err
	}

	linkedAccountID, err := uuid.Parse(savedLinkedAccount.GetID())
	if err != nil {
		return nil, fmt.Errorf("error parsing linked account ID: %v", err)
	}

	return &LinkAccountCommandResponse{
		AccountID: linkedAccountID,
		Message:   "Utility account linked successfully",
	}, nil
}
