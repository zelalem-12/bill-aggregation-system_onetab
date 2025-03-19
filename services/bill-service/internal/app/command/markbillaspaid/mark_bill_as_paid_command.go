package markbillaspaid

import "github.com/google/uuid"

type MarkBillAsPaidQuery struct {
	BillID uuid.UUID
}

type MarkBillAsPaidQueryResponse struct {
	Message string
}
