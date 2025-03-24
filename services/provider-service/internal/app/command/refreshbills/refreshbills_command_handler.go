package refreshbills

import (
	"context"
	"log"
	"sync"

	"github.com/google/uuid"
	clientPort "github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/repo"
)

type RefreshBillsCommandHandler struct {
	providerRepo   repo.ProviderRepo
	userClient     clientPort.UserServiceClient
	providerClient clientPort.ProviderServiceClient
	billClient     clientPort.BillServiceClient
}

func NewRefreshBillsCommandHandler(
	providerRepo repo.ProviderRepo,
	userClient clientPort.UserServiceClient,
	providerClient clientPort.ProviderServiceClient,
	billClient clientPort.BillServiceClient,
) *RefreshBillsCommandHandler {
	return &RefreshBillsCommandHandler{
		providerRepo:   providerRepo,
		userClient:     userClient,
		providerClient: providerClient,
		billClient:     billClient,
	}
}

func (h *RefreshBillsCommandHandler) Handle(ctx context.Context, cmd *RefreshBillsCommand) (*RefreshBillsCommandResponse, error) {

	userDetails, err := h.userClient.GetUserDetail(cmd.UserID)
	if err != nil {
		log.Printf("Failed to fetch user details: %v", err)
		return nil, err
	}

	var wg sync.WaitGroup

	for _, account := range userDetails.LinkedAccounts {
		wg.Add(1)
		go func(acc *clientPort.LinkedAccount) {
			defer wg.Done()
			err := h.processLinkedAccountBills(ctx, cmd.UserID, acc)
			if err != nil {
				log.Printf("Failed to refresh bills for account %s: %v", acc.ID, err)
			}
		}(account)
	}

	wg.Wait()

	return &RefreshBillsCommandResponse{Message: "Bills refreshed successfully"}, nil
}

func (h *RefreshBillsCommandHandler) processLinkedAccountBills(ctx context.Context, userID uuid.UUID, account *clientPort.LinkedAccount) error {
	provider, err := h.providerRepo.FindByID(ctx, account.ProviderID)
	if err != nil {
		return err
	}

	response, err := h.providerClient.FetchBillsFromProvider(account, provider)
	if err != nil {
		return err
	}
	// UserID     uuid.UUID `json:"user_id" validate:"required"`
	// Amount     float64   `json:"amount" validate:"required,gt=0"`
	// DueDate    time.Time `json:"due_date" validate:"required"`
	// Status     string    `json:"status" validate:"required,oneof=pending paid"`
	// ProviderID uuid.UUID `json:"provider_id" validate:"required"`
	// PaidDate   time.Time `json:"paid_date"`

	bills := make([]*clientPort.CreateBillRequestDTO, 0)
	for _, bill := range response {
		bills = append(bills, &clientPort.CreateBillRequestDTO{
			UserID:     userID,
			Amount:     bill.Amount,
			DueDate:    bill.DueDate,
			Status:     bill.Status,
			ProviderID: account.ProviderID,
			PaidDate:   bill.PaidDate,
		})
	}

	for _, bill := range bills {
		_, err = h.billClient.CreateBill(userID, bill)
		if err != nil {
			return err
		}
	}
	return nil
}
