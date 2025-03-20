package overduebills

import "github.com/google/uuid"

type OverdueBill struct {
	BillID    uuid.UUID
	AmountDue float64
	DueDate   string
}

type GetOverdueBillsQuery struct {
	UserID uuid.UUID
}

type GetOverdueBillsQueryResponse struct {
	Bills []OverdueBill
}
