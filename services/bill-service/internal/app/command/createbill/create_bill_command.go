package createbill

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateBillCommand struct {
	Amount       float64   `json:"amount" validate:"required"`
	DueDate      time.Time `json:"due_date" validate:"required"`
	Status       string    `json:"status" validate:"required"`
	UserID       uuid.UUID `json:"user_id" validate:"required"`
	ProviderID   uuid.UUID `json:"provider_id" validate:"required"`
	ProviderName string    `json:"provider_name" validate:"required"`

	PaidDate time.Time `json:"paid_date"`
}

func (c *CreateBillCommand) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

type CreateBillCommandResponse struct {
	BillID uuid.UUID
}
