package repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/adapter/persistence/postgres/model"
	repoPort "github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/domain"
	"gorm.io/gorm"
)

type ProviderRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) repoPort.ProviderRepo {
	return &ProviderRepo{
		DB: DB,
	}
}

func (repo *ProviderRepo) Save(ctx context.Context, domainModel *domain.Provider) (*domain.Provider, error) {
	dataModel := &model.Provider{}

	err := dataModel.FromDomainModel(domainModel)
	if err != nil {
		fmt.Println("ProviderRepo : Error converting domain model to data model", err)
		return nil, err
	}

	err = repo.DB.WithContext(ctx).Save(dataModel).Error
	if err != nil {
		fmt.Println("ProviderRepo : Error saving user", err)
		return nil, err
	}

	provider := dataModel.ToDomainModel()

	return provider, nil
}

func (repo *ProviderRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.Provider, error) {
	dataModel := &model.Provider{}

	err := repo.DB.WithContext(ctx).Where("id = ?", id).First(dataModel).Error
	if err != nil {
		return nil, err
	}

	return dataModel.ToDomainModel(), nil

}

func (repo *ProviderRepo) FindAll(ctx context.Context) ([]*domain.Provider, error) {
	dataModels := []*model.Provider{}

	err := repo.DB.WithContext(ctx).Find(&dataModels).Error
	if err != nil {
		return nil, err
	}

	domainModels := make([]*domain.Provider, 0, len(dataModels))

	for _, dataModel := range dataModels {

		domainModels = append(domainModels, dataModel.ToDomainModel())
	}

	return domainModels, nil

}
