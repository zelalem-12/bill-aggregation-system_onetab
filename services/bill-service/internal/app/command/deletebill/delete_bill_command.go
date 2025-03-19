package deletebill

import "github.com/google/uuid"

type DeleteBillQuery struct {
	BillID uuid.UUID
}

type DeleteBillQueryResponse struct {
	Message string
}
