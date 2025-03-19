package markbillaspaid

import (
	"context"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
)

type MarkBillAsPaidQueryHandler struct {
	BillRepo repo.BillRepo
}

func NewMarkBillAsPaidQueryHandler(billRepo repo.BillRepo) *MarkBillAsPaidQueryHandler {
	return &MarkBillAsPaidQueryHandler{
		BillRepo: billRepo,
	}
}

func (h *MarkBillAsPaidQueryHandler) Handle(ctx context.Context, query *MarkBillAsPaidQuery) (*MarkBillAsPaidQueryResponse, error) {
	if err := h.BillRepo.MarkAsPaid(ctx, query.BillID); err != nil {
		return nil, err
	}
	return &MarkBillAsPaidQueryResponse{
		Message: "Bill marked as paid",
	}, nil
}
