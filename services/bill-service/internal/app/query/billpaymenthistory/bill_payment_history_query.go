package billpaymenthistory

import "github.com/google/uuid"

type PaymentHistory struct {
	BillID   uuid.UUID
	Amount   float64
	PaidDate string
	DueDate  string
}

type GetBillPaymentHistoryQuery struct {
	UserID uuid.UUID
}

type GetBillPaymentHistoryQueryResponse struct {
	History []PaymentHistory
}
