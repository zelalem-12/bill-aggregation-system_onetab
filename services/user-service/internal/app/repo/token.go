package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/domain"
)

type TokenRepo interface {
	Save(ctx context.Context, token *domain.Token) (*domain.Token, error)
	Find(ctx context.Context, userId uuid.UUID, token string) (*domain.Token, error)
	Delete(ctx context.Context, userId uuid.UUID, token string) error
	DeleteByUserID(ctx context.Context, userId uuid.UUID) error
	FindByUserID(ctx context.Context, id uuid.UUID) (*domain.Token, error)
}
