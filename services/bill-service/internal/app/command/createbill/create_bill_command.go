package createbill

import (
	"time"

	"github.com/google/uuid"
)

type CreateBillCommand struct {
	Amount     float64
	DueDate    time.Time
	Status     string
	UserID     uuid.UUID
	ProviderID uuid.UUID
}

type CreateBillCommandResponse struct {
	BillID uuid.UUID
}
