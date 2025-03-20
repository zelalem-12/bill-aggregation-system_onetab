package billsummary

import "github.com/google/uuid"

type GetBillSummaryQuery struct {
	UserID uuid.UUID
}

type GetBillSummaryQueryResponse struct {
	TotalAmountDue float64
	TotalPaid      float64
	TotalOverdue   float64
}
