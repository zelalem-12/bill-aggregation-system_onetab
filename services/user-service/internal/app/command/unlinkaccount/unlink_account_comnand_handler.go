package unlinkaccount

import (
	"context"
	"fmt"

	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/util"
)

type UnlinkAccountCommandHandler struct {
	config            *config.Config
	linkedAccountRepo repo.LinkedAccountRepo
	billServiceClient client.BillServiceClient
}

func NewUnlinkAccountCommandHandler(
	config *config.Config,
	linkedAccountRepo repo.LinkedAccountRepo,
	billServiceClient client.BillServiceClient,
) *UnlinkAccountCommandHandler {
	return &UnlinkAccountCommandHandler{
		config:            config,
		linkedAccountRepo: linkedAccountRepo,
		billServiceClient: billServiceClient,
	}
}

func (h *UnlinkAccountCommandHandler) Handle(ctx context.Context, command *UnlinkAccountCommand) (*UnlinkAccountCommandResponse, error) {
	if err := command.Validate(); err != nil {
		return nil, err
	}

	account, err := h.linkedAccountRepo.FindByID(ctx, command.AccountID)
	if err != nil {
		return nil, fmt.Errorf("error finding account: %w", err)
	}

	if account == nil {
		return nil, fmt.Errorf("linked account with ID %s not found", command.AccountID)
	}

	if err := h.linkedAccountRepo.Delete(ctx, command.AccountID); err != nil {
		return nil, fmt.Errorf("error unlinking account: %w", err)
	}

	userID, err := util.ToUUID(account.GetUserID())
	if err != nil {
		return nil, fmt.Errorf("error converting user ID: %w", err)
	}

	providerID, err := util.ToUUID(account.GetProviderID())
	if err != nil {
		return nil, fmt.Errorf("error converting provider ID: %w", err)
	}

	_, err = h.billServiceClient.RemoveUnlinkedProviderBills(userID, providerID)
	if err != nil {
		return nil, fmt.Errorf("error deleting account from bill service: %w", err)
	}

	return &UnlinkAccountCommandResponse{
		AccountID: command.AccountID,
		Message:   "Utility account unlinked successfully",
	}, nil
}
