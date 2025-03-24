package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/command/createbill"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/command/deletebillsbyprovider"
)

type GetAggregatedBillsRequest struct {
	UserID uuid.UUID `param:"user_id" validate:"required"`
}

func (r *GetAggregatedBillsRequest) Validate() error {
	return validate.Struct(r)
}

type GetBillsByProviderRequest struct {
	ProviderName string `param:"provider_name" validate:"required"`
}

func (r *GetBillsByProviderRequest) Validate() error {
	return validate.Struct(r)
}

type GetBillsByProviderIdRequest struct {
	ProviderId string `param:"provider_id" validate:"required"`
}

func (r *GetBillsByProviderIdRequest) Validate() error {
	return validate.Struct(r)
}

type RefreshBillsRequest struct {
	UserID uuid.UUID `param:"user_id" validate:"required"`
}

func (r *RefreshBillsRequest) Validate() error {
	return validate.Struct(r)
}

type MarkBillAsPaidRequest struct {
	BillID uuid.UUID `param:"bill_id" validate:"required"`
}

func (r *MarkBillAsPaidRequest) Validate() error {
	return validate.Struct(r)
}

type DeleteBillRequest struct {
	BillID uuid.UUID `param:"bill_id" validate:"required"`
}

func (r *DeleteBillRequest) Validate() error {
	return validate.Struct(r)
}

type CreateBillRequest struct {
	UserID     uuid.UUID `header:"user_id" validate:"required"`
	Amount     float64   `json:"amount" validate:"required,gt=0"`
	DueDate    time.Time `json:"due_date" validate:"required"`
	Status     string    `json:"status" validate:"required,oneof=pending paid"`
	ProviderID uuid.UUID `json:"provider_id" validate:"required"`
	PaidDate   time.Time `json:"paid_date"`
}

func (r *CreateBillRequest) Validate() error {
	return validate.Struct(r)
}

func (r *CreateBillRequest) ToCommand() (*createbill.CreateBillCommand, error) {

	return &createbill.CreateBillCommand{
		UserID:     r.UserID,
		Amount:     r.Amount,
		DueDate:    r.DueDate,
		Status:     r.Status,
		ProviderID: r.ProviderID,
		PaidDate:   r.PaidDate,
	}, nil
}

type DeleteBillsByProviderRequest struct {
	ProviderID uuid.UUID `param:"provider_id" validate:"required"`
	UserID     uuid.UUID `header:"user_id" validate:"required"`
}

func (r *DeleteBillsByProviderRequest) Validate() error {
	return validate.Struct(r)
}

func (r *DeleteBillsByProviderRequest) ToCommand() *deletebillsbyprovider.DeleteBillsByProviderCommand {
	return &deletebillsbyprovider.DeleteBillsByProviderCommand{
		ProviderID: r.ProviderID,
		UserID:     r.UserID,
	}
}
