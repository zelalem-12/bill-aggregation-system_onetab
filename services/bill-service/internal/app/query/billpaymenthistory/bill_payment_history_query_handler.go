package billpaymenthistory

import (
	"context"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/service"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
)

type GetBillPaymentHistoryQueryHandler struct {
	BillRepo repo.BillRepo
}

func NewGetBillPaymentHistoryQueryHandler(billRepo repo.BillRepo) *GetBillPaymentHistoryQueryHandler {
	return &GetBillPaymentHistoryQueryHandler{BillRepo: billRepo}
}

func (h *GetBillPaymentHistoryQueryHandler) Handle(ctx context.Context, query *GetBillPaymentHistoryQuery) (*GetBillPaymentHistoryQueryResponse, error) {
	bills, err := h.BillRepo.FindByUserID(ctx, query.UserID)
	if err != nil {
		return nil, err
	}
	var history []PaymentHistory
	for _, bill := range bills {
		if bill.GetStatus() == domain.PAID {
			paidDate := ""
			if !bill.GetPaidDate().IsZero() {
				paidDate = bill.GetPaidDate().Format("2006-01-02")
			}
			dueDate := ""
			if !bill.GetDueDate().IsZero() {
				dueDate = bill.GetDueDate().Format("2006-01-02")
			}
			billID, _ := service.ToUUID(bill.GetID())
			history = append(history, PaymentHistory{
				BillID:   billID,
				Amount:   bill.GetAmount(),
				PaidDate: paidDate,
				DueDate:  dueDate,
			})
		}
	}
	return &GetBillPaymentHistoryQueryResponse{History: history}, nil
}
