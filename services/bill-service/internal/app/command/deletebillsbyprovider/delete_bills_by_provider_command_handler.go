package deletebillsbyprovider

import (
	"context"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
)

type DeleteBillsByProviderCommandHandler struct {
	BillRepo repo.BillRepo
}

func NewDeleteBillsByProviderCommandHandler(billRepo repo.BillRepo) *DeleteBillsByProviderCommandHandler {
	return &DeleteBillsByProviderCommandHandler{
		BillRepo: billRepo,
	}
}

func (h *DeleteBillsByProviderCommandHandler) Handle(ctx context.Context, cmd *DeleteBillsByProviderCommand) (*DeleteBillsByProviderCommandResponse, error) {
	if err := h.BillRepo.DeleteByProvider(ctx, cmd.UserID, cmd.ProviderID); err != nil {
		return nil, err
	}

	return &DeleteBillsByProviderCommandResponse{
		Message: "Bills for the provider deleted successfully",
	}, nil
}
