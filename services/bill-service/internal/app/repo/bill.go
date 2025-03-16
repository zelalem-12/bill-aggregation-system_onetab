package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
)

type BillRepo interface {
	Save(ctx context.Context, user *domain.Bill) (*domain.Bill, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Bill, error)
	FindAll(ctx context.Context) ([]*domain.Bill, error)
}
