package aggregatedbills

import (
	"context"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/util"
)

type GetAggregatedBillsQueryHandler struct {
	BillRepo repo.BillRepo
}

func NewGetAggregatedBillsQueryHandler(billRepo repo.BillRepo) *GetAggregatedBillsQueryHandler {
	return &GetAggregatedBillsQueryHandler{
		BillRepo: billRepo,
	}
}

func (h *GetAggregatedBillsQueryHandler) Handle(ctx context.Context, query *GetAggregatedBillsQuery) (*GetAggregatedBillsQueryResponse, error) {

	bills, err := h.BillRepo.FindByUserID(ctx, query.UserID)
	if err != nil {
		return nil, err
	}

	var totalDue float64
	var billResponses []Bill

	for _, bill := range bills {

		if bill.GetStatus() == domain.PENDING {
			totalDue += bill.GetAmount()
		}

		billID, err := util.ToUUID(bill.GetID())
		if err != nil {
			return nil, err
		}

		billResponses = append(billResponses, Bill{
			ID:           billID,
			ProviderName: bill.GetProviderName(),
			Amount:       bill.GetAmount(),
			DueDate:      bill.GetDueDate().Format("2006-01-02"),
			Status:       bill.GetStatus().String(),
		})
	}

	return &GetAggregatedBillsQueryResponse{
		TotalDue: totalDue,
		Bills:    billResponses,
	}, nil
}
