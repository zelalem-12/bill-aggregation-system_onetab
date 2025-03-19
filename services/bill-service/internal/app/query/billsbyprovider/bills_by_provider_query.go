package billsbyprovider

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type GetBillsByProviderQuery struct {
	UserID       uuid.UUID `json:"user_id" validate:"required"`
	ProviderID   *uuid.UUID
	ProviderName *string
}

func (q *GetBillsByProviderQuery) Validate() error {
	validate := validator.New()
	return validate.Struct(q)
}

type Bill struct {
	ID      uuid.UUID
	Amount  float64
	DueDate string // e.g., "2006-01-02"
	Status  string
}

type GetBillsByProviderQueryResponse struct {
	Provider string
	Bills    []Bill
}
