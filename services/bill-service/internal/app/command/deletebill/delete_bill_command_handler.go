package deletebill

import (
	"context"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
)

type DeleteBillQueryHandler struct {
	BillRepo repo.BillRepo
}

func NewDeleteBillQueryHandler(billRepo repo.BillRepo) *DeleteBillQueryHandler {
	return &DeleteBillQueryHandler{
		BillRepo: billRepo,
	}
}

func (h *DeleteBillQueryHandler) Handle(ctx context.Context, query *DeleteBillQuery) (*DeleteBillQueryResponse, error) {
	if err := h.BillRepo.Delete(ctx, query.BillID); err != nil {
		return nil, err
	}
	return &DeleteBillQueryResponse{
		Message: "Bill deleted successfully",
	}, nil
}
