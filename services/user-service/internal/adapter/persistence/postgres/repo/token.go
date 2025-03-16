package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/zelalem-12/onetab/internal/adapter/persistence/postgres/model"
	repoPort "github.com/zelalem-12/onetab/internal/app/repo"
	"github.com/zelalem-12/onetab/internal/domain"
	"gorm.io/gorm"
)

type TokenRepo struct {
	DB *gorm.DB
}

func NewToken(DB *gorm.DB) repoPort.TokenRepo {
	return &TokenRepo{
		DB: DB,
	}
}

func (repo *TokenRepo) Save(ctx context.Context, token *domain.Token) (*domain.Token, error) {
	dataModel := &model.Token{}

	dataModel.FromDomainModel(token)

	err := repo.DB.WithContext(ctx).Save(dataModel).Error
	if err != nil {
		return nil, err
	}

	return dataModel.ToDomainModel(), nil
}

func (repo *TokenRepo) Find(ctx context.Context, userId uuid.UUID, token string) (*domain.Token, error) {
	dataModel := &model.Token{}

	err := repo.DB.WithContext(ctx).Where("user_id = ? AND token = ?", userId, token).First(dataModel).Error
	if err != nil {
		return nil, err
	}

	return dataModel.ToDomainModel(), nil
}

func (repo *TokenRepo) Delete(ctx context.Context, userId uuid.UUID, token string) error {

	return repo.DB.WithContext(ctx).Where("user_id = ? AND token = ?", userId, token).Delete(&model.Token{}).Error

}

func (repo *TokenRepo) DeleteByUserID(ctx context.Context, userId uuid.UUID) error {

	return repo.DB.WithContext(ctx).Where("user_id = ?", userId).Delete(&model.Token{}).Error
}
