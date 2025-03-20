package spendingtrends

import "github.com/google/uuid"

type SpendingTrend struct {
	Month      string  `json:"month"`
	TotalSpent float64 `json:"total_spent"`
}

type GetMonthlySpendingTrendsQuery struct {
	UserID uuid.UUID
}

type GetMonthlySpendingTrendsQueryResponse struct {
	Trends []SpendingTrend `json:"trends"`
}
