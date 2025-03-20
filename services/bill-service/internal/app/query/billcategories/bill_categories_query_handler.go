package billcategories

import (
	"context"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
)

type GetCategorySpendingQueryHandler struct {
	BillRepo repo.BillRepo
}

func NewGetCategorySpendingQueryHandler(billRepo repo.BillRepo) *GetCategorySpendingQueryHandler {
	return &GetCategorySpendingQueryHandler{BillRepo: billRepo}
}

func (h *GetCategorySpendingQueryHandler) Handle(ctx context.Context, query *GetCategorySpendingQuery) (*GetCategorySpendingQueryResponse, error) {
	bills, err := h.BillRepo.FindByUserID(ctx, query.UserID)
	if err != nil {
		return nil, err
	}

	categorySpending := make(map[string]float64)
	for _, bill := range bills {
		if bill.GetStatus() == domain.PAID {
			categorySpending[bill.GetProviderName()] += bill.GetAmount()
		}
	}

	var categories []CategorySpending
	for cat, total := range categorySpending {
		categories = append(categories, CategorySpending{
			Category:   cat,
			TotalSpent: total,
		})
	}

	return &GetCategorySpendingQueryResponse{
		Categories: categories,
	}, nil
}
