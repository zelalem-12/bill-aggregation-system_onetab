package response

import (
	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/command/deletebillsbyprovider"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/aggregatedbills"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/billsbyprovider"
)

type BillResponse struct {
	ID       uuid.UUID `json:"id"`
	Provider string    `json:"provider"`
	Amount   float64   `json:"amount"`
	DueDate  string    `json:"due_date"`
	Status   string    `json:"status"`
}

type GetAggregatedBillsResponse struct {
	TotalDue float64        `json:"total_due"`
	Bills    []BillResponse `json:"bills"`
}

func (r *GetAggregatedBillsResponse) FromQueryResponse(queryResponse *aggregatedbills.GetAggregatedBillsQueryResponse) {
	r.TotalDue = queryResponse.TotalDue
	r.Bills = make([]BillResponse, len(queryResponse.Bills))
	for i, bill := range queryResponse.Bills {
		r.Bills[i] = BillResponse{
			ID:       bill.ID,
			Provider: bill.ProviderName,
			Amount:   bill.Amount,
			DueDate:  bill.DueDate,
			Status:   bill.Status,
		}
	}
}

type GetBillsByProviderResponse struct {
	Provider string         `json:"provider"`
	Bills    []BillResponse `json:"bills"`
}

func (r *GetBillsByProviderResponse) FromQueryResponse(queryResponse *billsbyprovider.GetBillsByProviderQueryResponse) {

	r.Provider = queryResponse.Provider
	r.Bills = make([]BillResponse, len(queryResponse.Bills))
	for i, bill := range queryResponse.Bills {
		r.Bills[i] = BillResponse{
			ID:      bill.ID,
			Amount:  bill.Amount,
			DueDate: bill.DueDate,
			Status:  bill.Status,
		}
	}
}

type CreateBillResponse struct {
	ID      uuid.UUID `json:"id"`
	Message string    `json:"message"`
}

type DeleteBillsByProviderResponse struct {
	Message string `json:"message"`
}

func (r *DeleteBillsByProviderResponse) FromCommandResponse(cmdResp *deletebillsbyprovider.DeleteBillsByProviderCommandResponse) {
	r.Message = cmdResp.Message
}
