package createbill

import (
	"context"
	"fmt"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/service"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
)

type CreateBillCommandHandler struct {
	BillRepo repo.BillRepo
}

func NewCreateBillCommandHandler(billRepo repo.BillRepo) *CreateBillCommandHandler {
	return &CreateBillCommandHandler{
		BillRepo: billRepo,
	}
}

func (h *CreateBillCommandHandler) Handle(ctx context.Context, cmd *CreateBillCommand) (*CreateBillCommandResponse, error) {

	bill := domain.NewBill(cmd.UserID.String(), cmd.ProviderID.String(), cmd.Amount, domain.BillStatus(cmd.Status), cmd.DueDate)

	persistedBill, err := h.BillRepo.Save(ctx, bill)
	if err != nil {
		return nil, fmt.Errorf("error saving bill: %v", err)
	}

	billID, err := service.ToUUID(persistedBill.GetID())
	if err != nil {
		return nil, fmt.Errorf("error parsing linked account ID: %v", err)
	}

	return &CreateBillCommandResponse{
		BillID: billID,
	}, nil
}
