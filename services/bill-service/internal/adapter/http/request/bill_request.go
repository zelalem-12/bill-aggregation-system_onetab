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
	Amount     float64   `json:"amount" validate:"required,gt=0"`
	DueDate    string    `json:"due_date" validate:"required,datetime=2006-01-02"`
	Status     string    `json:"status" validate:"required,oneof=paid unpaid"`
	ProviderID uuid.UUID `json:"provider_id" validate:"required"`
}

func (r *CreateBillRequest) Validate() error {
	return validate.Struct(r)
}

func (r *CreateBillRequest) ToCommand() (*createbill.CreateBillCommand, error) {
	dt, err := time.Parse("2006-01-02", r.DueDate)
	if err != nil {
		return nil, err
	}
	return &createbill.CreateBillCommand{
		Amount:     r.Amount,
		DueDate:    dt,
		Status:     r.Status,
		ProviderID: r.ProviderID,
	}, nil
}

type DeleteBillsByProviderRequest struct {
	ProviderID uuid.UUID `param:"provider_id" validate:"required"`
}

func (r *DeleteBillsByProviderRequest) Validate() error {
	return validate.Struct(r)
}

func (r *DeleteBillsByProviderRequest) ToCommand() *deletebillsbyprovider.DeleteBillsByProviderCommand {
	return &deletebillsbyprovider.DeleteBillsByProviderCommand{
		ProviderID: r.ProviderID,
	}
}
