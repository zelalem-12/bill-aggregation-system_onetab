package aggregatedbills

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type GetAggregatedBillsQuery struct {
	UserID uuid.UUID
}

func (q *GetAggregatedBillsQuery) Validate() error {
	validate := validator.New()
	return validate.Struct(q)
}

type Bill struct {
	ID           uuid.UUID
	ProviderName string
	Amount       float64
	DueDate      string // e.g., "2006-01-02"
	Status       string
}

type GetAggregatedBillsQueryResponse struct {
	TotalDue float64
	Bills    []Bill
}
