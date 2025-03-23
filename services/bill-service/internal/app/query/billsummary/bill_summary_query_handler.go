package billsummary

import (
	"context"
	"time"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
)

type GetBillSummaryQueryHandler struct {
	BillRepo repo.BillRepo
}

func NewGetBillSummaryQueryHandler(billRepo repo.BillRepo) *GetBillSummaryQueryHandler {
	return &GetBillSummaryQueryHandler{BillRepo: billRepo}
}

func (h *GetBillSummaryQueryHandler) Handle(ctx context.Context, query *GetBillSummaryQuery) (*GetBillSummaryQueryResponse, error) {
	bills, err := h.BillRepo.FindByUserID(ctx, query.UserID)
	if err != nil {
		return nil, err
	}

	var totalOutstanding, totalOverdue, totalPaid float64
	currentTime := time.Now()

	for _, bill := range bills {
		switch bill.GetStatus() {
		case domain.PENDING:
			if bill.GetDueDate().Before(currentTime) {
				totalOverdue += bill.GetAmount()
			} else {
				totalOutstanding += bill.GetAmount()
			}
		case domain.PAID:
			totalPaid += bill.GetAmount()
		}
	}

	return &GetBillSummaryQueryResponse{
		TotalAmountDue: totalOutstanding + totalOverdue, // All unpaid bills
		TotalOverdue:   totalOverdue,                    // Overdue bills
		TotalPaid:      totalPaid,                       // Paid bills
	}, nil
}
