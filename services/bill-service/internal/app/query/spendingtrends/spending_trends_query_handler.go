package spendingtrends

import (
	"context"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
)

type GetMonthlySpendingTrendsQueryHandler struct {
	BillRepo repo.BillRepo
}

func NewGetMonthlySpendingTrendsQueryHandler(billRepo repo.BillRepo) *GetMonthlySpendingTrendsQueryHandler {
	return &GetMonthlySpendingTrendsQueryHandler{BillRepo: billRepo}
}

func (h *GetMonthlySpendingTrendsQueryHandler) Handle(ctx context.Context, query *GetMonthlySpendingTrendsQuery) (*GetMonthlySpendingTrendsQueryResponse, error) {
	bills, err := h.BillRepo.FindByUserID(ctx, query.UserID)
	if err != nil {
		return nil, err
	}
	monthlySpending := make(map[string]float64)
	for _, bill := range bills {
		if bill.GetStatus() == domain.PAID {
			month := bill.GetDueDate().Format("2006-01")
			monthlySpending[month] += bill.GetAmount()
		}
	}
	var trends []SpendingTrend
	for month, total := range monthlySpending {
		trends = append(trends, SpendingTrend{
			Month:      month,
			TotalSpent: total,
		})
	}
	return &GetMonthlySpendingTrendsQueryResponse{
		Trends: trends,
	}, nil
}
