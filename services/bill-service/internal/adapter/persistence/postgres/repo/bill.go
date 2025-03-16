package repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/adapter/persistence/postgres/model"
	repoPort "github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
	"gorm.io/gorm"
)

type BillRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) repoPort.BillRepo {
	return &BillRepo{
		DB: DB,
	}
}

func (repo *BillRepo) Save(ctx context.Context, domainModel *domain.Bill) (*domain.Bill, error) {
	dataModel := &model.Bill{}

	err := dataModel.FromDomainModel(domainModel)
	if err != nil {
		fmt.Println("BilRepo : Error converting domain model to data model", err)
		return nil, err
	}

	err = repo.DB.WithContext(ctx).Save(dataModel).Error
	if err != nil {
		fmt.Println("BilRepo : Error saving user", err)
		return nil, err
	}

	provider := dataModel.ToDomainModel()

	return provider, nil
}

func (repo *BillRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.Bill, error) {
	dataModel := &model.Bill{}

	err := repo.DB.WithContext(ctx).Where("id = ?", id).First(dataModel).Error
	if err != nil {
		return nil, err
	}

	return dataModel.ToDomainModel(), nil

}

func (repo *BillRepo) FindAll(ctx context.Context) ([]*domain.Bill, error) {
	dataModels := []*model.Bill{}

	err := repo.DB.WithContext(ctx).Find(&dataModels).Error
	if err != nil {
		return nil, err
	}

	domainModels := make([]*domain.Bill, 0, len(dataModels))

	for _, dataModel := range dataModels {

		domainModels = append(domainModels, dataModel.ToDomainModel())
	}

	return domainModels, nil

}
