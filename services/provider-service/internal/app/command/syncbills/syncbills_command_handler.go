package syncbills

import (
	"context"
	"log"
	"sync"

	"github.com/google/uuid"
	clientPort "github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/repo"
)

type SyncAllBillsCommandHandler struct {
	providerRepo   repo.ProviderRepo
	userClient     clientPort.UserServiceClient
	providerClient clientPort.ProviderServiceClient
	billClient     clientPort.BillServiceClient
}

func NewSyncAllBillsCommandHandler(
	providerRepo repo.ProviderRepo,
	userClient clientPort.UserServiceClient,
	providerClient clientPort.ProviderServiceClient,
	billClient clientPort.BillServiceClient,
) *SyncAllBillsCommandHandler {
	return &SyncAllBillsCommandHandler{
		providerRepo:   providerRepo,
		userClient:     userClient,
		providerClient: providerClient,
		billClient:     billClient,
	}
}

func (h *SyncAllBillsCommandHandler) Handle(ctx context.Context, cmd *SyncAllBillsCommand) (*SyncAllBillsCommandResponse, error) {
	response, err := h.userClient.GetUsers()
	if err != nil {
		log.Printf("Failed to fetch users: %v", err)
		return nil, err
	}

	var wg sync.WaitGroup
	for _, user := range response.Users {
		wg.Add(1)
		go func(user *clientPort.UserDetail) {
			defer wg.Done()
			h.processUserBills(ctx, user)
		}(user)
	}
	wg.Wait()

	return &SyncAllBillsCommandResponse{Message: "All users' bills synced successfully"}, nil
}

func (h *SyncAllBillsCommandHandler) processUserBills(ctx context.Context, user *clientPort.UserDetail) {
	var wg sync.WaitGroup
	for _, account := range user.LinkedAccounts {
		wg.Add(1)
		go func(acc *clientPort.LinkedAccount) {
			defer wg.Done()
			if err := h.processLinkedAccountBills(ctx, user.ID, acc); err != nil {
				log.Printf("Failed to sync bills for user %s, account %s: %v", user.ID, acc.ID, err)
			}
		}(account)
	}
	wg.Wait()
}

func (h *SyncAllBillsCommandHandler) processLinkedAccountBills(ctx context.Context, userID uuid.UUID, account *clientPort.LinkedAccount) error {
	provider, err := h.providerRepo.FindByID(ctx, account.ProviderID)
	if err != nil {
		return err
	}

	response, err := h.providerClient.FetchBillsFromProvider(account, provider)
	if err != nil {
		return err
	}

	bills := make([]*clientPort.CreateBillRequestDTO, len(response))
	for i, bill := range response {
		bills[i] = &clientPort.CreateBillRequestDTO{
			UserID:     userID,
			Amount:     bill.Amount,
			DueDate:    bill.DueDate,
			Status:     bill.Status,
			ProviderID: account.ProviderID,
			PaidDate:   bill.PaidDate,
		}
	}

	for _, bill := range bills {
		_, err = h.billClient.CreateBill(userID, bill)
		if err != nil {
			return err
		}
	}
	return nil
}
