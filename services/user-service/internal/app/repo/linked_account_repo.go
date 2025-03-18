package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/domain"
)

type LinkedAccountRepo interface {
	Save(ctx context.Context, account *domain.LinkedAccount) (*domain.LinkedAccount, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.LinkedAccount, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.LinkedAccount, error)
	FindByProvider(ctx context.Context, userID uuid.UUID, provider string) (*domain.LinkedAccount, error)
	Update(ctx context.Context, account *domain.LinkedAccount) (*domain.LinkedAccount, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
}
