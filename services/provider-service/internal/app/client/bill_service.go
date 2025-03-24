package client

import (
	"time"

	"github.com/google/uuid"
)

type CreateBillRequestDTO struct {
	UserID     uuid.UUID `json:"user_id" validate:"required"`
	Amount     float64   `json:"amount" validate:"required,gt=0"`
	DueDate    time.Time `json:"due_date" validate:"required"`
	Status     string    `json:"status" validate:"required,oneof=pending paid"`
	ProviderID uuid.UUID `json:"provider_id" validate:"required"`
	PaidDate   time.Time `json:"paid_date"`
}

type ProviderBillResponse struct {
	Amount   float64   `json:"amount"`
	DueDate  time.Time `json:"due_date"`
	Status   string    `json:"status"`
	PaidDate time.Time `json:"paid_date"`
}

type CreateBillResponse struct {
	ID      uuid.UUID `json:"id"`
	Message string    `json:"message"`
}

type BillServiceClient interface {
	CreateBill(userId uuid.UUID, bill *CreateBillRequestDTO) (*CreateBillResponse, error)
}
