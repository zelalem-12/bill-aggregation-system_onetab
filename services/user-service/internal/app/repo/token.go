package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/zelalem-12/onetab/internal/domain"
)

type TokenRepo interface {
	Save(ctx context.Context, token *domain.Token) (*domain.Token, error)
	Find(ctx context.Context, userId uuid.UUID, token string) (*domain.Token, error)
	Delete(ctx context.Context, userId uuid.UUID, token string) error
	DeleteByUserID(ctx context.Context, userId uuid.UUID) error
}
