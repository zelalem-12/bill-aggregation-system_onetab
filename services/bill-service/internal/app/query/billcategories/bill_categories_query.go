package billcategories

import "github.com/google/uuid"

type CategorySpending struct {
	Category   string  `json:"category"`
	TotalSpent float64 `json:"total_spent"`
}

type GetCategorySpendingQuery struct {
	UserID uuid.UUID
}

type GetCategorySpendingQueryResponse struct {
	Categories []CategorySpending `json:"categories"`
}
