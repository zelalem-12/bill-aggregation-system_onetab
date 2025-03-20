package overduebills

import (
	"context"
	"time"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/service"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
)

type GetOverdueBillsQueryHandler struct {
	BillRepo repo.BillRepo
}

func NewGetOverdueBillsQueryHandler(billRepo repo.BillRepo) *GetOverdueBillsQueryHandler {
	return &GetOverdueBillsQueryHandler{BillRepo: billRepo}
}

func (h *GetOverdueBillsQueryHandler) Handle(ctx context.Context, query *GetOverdueBillsQuery) (*GetOverdueBillsQueryResponse, error) {
	bills, err := h.BillRepo.FindByUserID(ctx, query.UserID)
	if err != nil {
		return nil, err
	}
	var overdue []OverdueBill
	now := time.Now()
	for _, bill := range bills {
		if bill.GetDueDate().Before(now) && bill.GetStatus() == domain.UNPAID {
			billID, _ := service.ToUUID(bill.GetID())
			overdue = append(overdue, OverdueBill{
				BillID:    billID,
				AmountDue: bill.GetAmount(),
				DueDate:   bill.GetDueDate().Format("2006-01-02"),
			})
		}
	}
	return &GetOverdueBillsQueryResponse{Bills: overdue}, nil
}
