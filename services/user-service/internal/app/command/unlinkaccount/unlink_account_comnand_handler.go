package unlinkaccount

import (
	"context"
	"fmt"

	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
)

type UnlinkAccountCommandHandler struct {
	config            *config.Config
	linkedAccountRepo repo.LinkedAccountRepo
}

func NewUnlinkAccountCommandHandler(
	config *config.Config,
	linkedAccountRepo repo.LinkedAccountRepo,
) *UnlinkAccountCommandHandler {
	return &UnlinkAccountCommandHandler{
		config:            config,
		linkedAccountRepo: linkedAccountRepo,
	}
}

func (h *UnlinkAccountCommandHandler) Handle(ctx context.Context, command *UnlinkAccountCommand) (*UnlinkAccountCommandResponse, error) {
	if err := command.Validate(); err != nil {
		return nil, err
	}

	exists, err := h.linkedAccountRepo.Exists(ctx, command.AccountID)
	if err != nil {
		return nil, fmt.Errorf("error checking linked account: %w", err)
	}

	if !exists {
		return nil, fmt.Errorf("linked account with ID %s not found", command.AccountID)
	}

	if err := h.linkedAccountRepo.Delete(ctx, command.AccountID); err != nil {
		return nil, fmt.Errorf("error unlinking account: %w", err)
	}

	return &UnlinkAccountCommandResponse{
		AccountID: command.AccountID,
		Message:   "Utility account unlinked successfully",
	}, nil
}
