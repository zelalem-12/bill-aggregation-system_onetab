package repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/persistence/postgres/model"
	repoPort "github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/domain"
	"gorm.io/gorm"
)

type LinkedAccountRepo struct {
	DB *gorm.DB
}

func NewLinkedAccountRepo(DB *gorm.DB) repoPort.LinkedAccountRepo {
	return &LinkedAccountRepo{
		DB: DB,
	}
}

func (repo *LinkedAccountRepo) Save(ctx context.Context, domainModel *domain.LinkedAccount) (*domain.LinkedAccount, error) {
	dataModel := &model.LinkedAccount{}

	err := dataModel.FromDomainModel(domainModel)
	if err != nil {
		fmt.Println("LinkedAccountRepo: Error converting domain model to data model", err)
		return nil, err
	}

	err = repo.DB.WithContext(ctx).Save(dataModel).Error
	if err != nil {
		fmt.Println("LinkedAccountRepo: Error saving linked account", err)
		return nil, err
	}

	return dataModel.ToDomainModel(), nil
}

func (repo *LinkedAccountRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.LinkedAccount, error) {
	dataModel := &model.LinkedAccount{}

	err := repo.DB.WithContext(ctx).Where("id = ?", id).First(dataModel).Error
	if err != nil {
		return nil, err
	}

	return dataModel.ToDomainModel(), nil
}

func (repo *LinkedAccountRepo) FindByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.LinkedAccount, error) {
	dataModels := []*model.LinkedAccount{}

	err := repo.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&dataModels).Error
	if err != nil {
		return nil, err
	}

	domainModels := make([]*domain.LinkedAccount, 0, len(dataModels))

	for _, dataModel := range dataModels {
		domainModels = append(domainModels, dataModel.ToDomainModel())
	}

	return domainModels, nil
}

func (repo *LinkedAccountRepo) FindByProvider(ctx context.Context, userID uuid.UUID, provider string) (*domain.LinkedAccount, error) {
	dataModel := &model.LinkedAccount{}

	err := repo.DB.WithContext(ctx).
		Where("user_id = ? AND provider = ?", userID, provider).
		First(dataModel).Error

	if err != nil {
		return nil, err
	}

	return dataModel.ToDomainModel(), nil
}

func (repo *LinkedAccountRepo) Update(ctx context.Context, domainModel *domain.LinkedAccount) (*domain.LinkedAccount, error) {
	dataModel := &model.LinkedAccount{}

	err := dataModel.FromDomainModel(domainModel)
	if err != nil {
		return nil, err
	}

	err = repo.DB.WithContext(ctx).Save(dataModel).Error
	if err != nil {
		return nil, err
	}

	return dataModel.ToDomainModel(), nil
}

func (repo *LinkedAccountRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return repo.DB.WithContext(ctx).Where("id = ?", id).Delete(&model.LinkedAccount{}).Error
}

func (repo *LinkedAccountRepo) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	return repo.DB.WithContext(ctx).Where("user_id = ?", userID).Delete(&model.LinkedAccount{}).Error
}

func (repo *LinkedAccountRepo) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64

	err := repo.DB.WithContext(ctx).Model(&model.LinkedAccount{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
