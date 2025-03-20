package response

import (
	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/command/deletebillsbyprovider"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/aggregatedbills"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/billcategories"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/billpaymenthistory"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/billsbyprovider"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/billsummary"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/overduebills"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/spendingtrends"
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

type OverdueBillResponse struct {
	BillID    string  `json:"bill_id"`
	AmountDue float64 `json:"amount_due"`
	DueDate   string  `json:"due_date"`
}

type GetOverdueBillsResponse struct {
	Bills []OverdueBillResponse `json:"bills"`
}

func (r *GetOverdueBillsResponse) FromQueryResponse(queryResp *overduebills.GetOverdueBillsQueryResponse) {
	r.Bills = make([]OverdueBillResponse, len(queryResp.Bills))
	for i, bill := range queryResp.Bills {
		r.Bills[i] = OverdueBillResponse{
			BillID:    bill.BillID.String(),
			AmountDue: bill.AmountDue,
			DueDate:   bill.DueDate,
		}
	}
}

type CategorySpendingResponse struct {
	Category   string  `json:"category"`
	TotalSpent float64 `json:"total_spent"`
}

type GetCategorySpendingResponse struct {
	Categories []CategorySpendingResponse `json:"categories"`
}

func (r *GetCategorySpendingResponse) FromQueryResponse(queryResp *billcategories.GetCategorySpendingQueryResponse) {
	r.Categories = make([]CategorySpendingResponse, len(queryResp.Categories))
	for i, cs := range queryResp.Categories {
		r.Categories[i] = CategorySpendingResponse{
			Category:   cs.Category,
			TotalSpent: cs.TotalSpent,
		}
	}
}

type PaymentHistoryResponse struct {
	BillID   string  `json:"bill_id"`
	Amount   float64 `json:"amount"`
	PaidDate string  `json:"paid_date"`
	DueDate  string  `json:"due_date"`
}

type GetBillPaymentHistoryResponse struct {
	History []PaymentHistoryResponse `json:"history"`
}

func (r *GetBillPaymentHistoryResponse) FromQueryResponse(queryResp *billpaymenthistory.GetBillPaymentHistoryQueryResponse) {
	r.History = make([]PaymentHistoryResponse, len(queryResp.History))
	for i, h := range queryResp.History {
		r.History[i] = PaymentHistoryResponse{
			BillID:   h.BillID.String(),
			Amount:   h.Amount,
			PaidDate: h.PaidDate,
			DueDate:  h.DueDate,
		}
	}
}

type GetBillSummaryResponse struct {
	TotalAmountDue float64 `json:"total_amount_due"`
	TotalPaid      float64 `json:"total_paid"`
	TotalOverdue   float64 `json:"total_overdue"`
}

func (r *GetBillSummaryResponse) FromQueryResponse(queryResp *billsummary.GetBillSummaryQueryResponse) {
	r.TotalAmountDue = queryResp.TotalAmountDue
	r.TotalPaid = queryResp.TotalPaid
	r.TotalOverdue = queryResp.TotalOverdue
}

type TrendResponse struct {
	Month      string  `json:"month"`
	TotalSpent float64 `json:"total_spent"`
}

type GetMonthlySpendingTrendsResponse struct {
	Trends []TrendResponse `json:"trends"`
}

func (r *GetMonthlySpendingTrendsResponse) FromQueryResponse(queryResp *spendingtrends.GetMonthlySpendingTrendsQueryResponse) {
	r.Trends = make([]TrendResponse, len(queryResp.Trends))
	for i, trend := range queryResp.Trends {
		r.Trends[i] = TrendResponse{
			Month:      trend.Month,
			TotalSpent: trend.TotalSpent,
		}
	}
}
