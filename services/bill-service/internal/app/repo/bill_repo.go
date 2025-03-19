package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
)

type BillRepo interface {
	FindByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Bill, error)
	FindByProviderID(ctx context.Context, userID uuid.UUID, providerID uuid.UUID) ([]*domain.Bill, error)
	FindByProviderName(ctx context.Context, userID uuid.UUID, providerName string) ([]*domain.Bill, error)
	FindByID(ctx context.Context, billID uuid.UUID) (*domain.Bill, error)
	Save(ctx context.Context, bill *domain.Bill) (*domain.Bill, error)
	MarkAsPaid(ctx context.Context, billID uuid.UUID) error
	Delete(ctx context.Context, billID uuid.UUID) error
	Upsert(ctx context.Context, bills []*domain.Bill) error
	DeleteByProvider(ctx context.Context, userID uuid.UUID, providerID uuid.UUID) error
}
