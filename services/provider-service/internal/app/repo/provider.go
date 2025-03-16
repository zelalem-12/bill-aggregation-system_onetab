package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/domain"
)

type ProviderRepo interface {
	Save(ctx context.Context, user *domain.Provider) (*domain.Provider, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Provider, error)
	FindAll(ctx context.Context) ([]*domain.Provider, error)
}
