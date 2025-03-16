package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/zelalem-12/onetab/internal/domain"
)

type UserRepo interface {
	Save(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	FindAll(ctx context.Context) ([]*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}
