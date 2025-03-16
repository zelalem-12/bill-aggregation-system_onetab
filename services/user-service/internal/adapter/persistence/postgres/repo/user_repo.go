package repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zelalem-12/onetab/internal/adapter/persistence/postgres/model"
	repoPort "github.com/zelalem-12/onetab/internal/app/repo"
	"github.com/zelalem-12/onetab/internal/domain"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) repoPort.UserRepo {
	return &UserRepo{
		DB: DB,
	}
}

func (repo *UserRepo) Save(ctx context.Context, domainModel *domain.User) (*domain.User, error) {
	dataModel := &model.User{}

	err := dataModel.FromDomainModel(domainModel)
	if err != nil {
		fmt.Println("UserRepo : Error converting domain model to data model", err)
		return nil, err
	}

	err = repo.DB.WithContext(ctx).Save(dataModel).Error
	if err != nil {
		fmt.Println("UserRepo : Error saving user", err)
		return nil, err
	}

	user := dataModel.ToDomainModel()

	return user, nil
}

func (repo *UserRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	dataModel := &model.User{}

	err := repo.DB.WithContext(ctx).Where("id = ?", id).First(dataModel).Error
	if err != nil {
		return nil, err
	}

	return dataModel.ToDomainModel(), nil

}

func (repo *UserRepo) FindAll(ctx context.Context) ([]*domain.User, error) {
	dataModels := []*model.User{}

	err := repo.DB.WithContext(ctx).Find(&dataModels).Error
	if err != nil {
		return nil, err
	}

	domainModels := make([]*domain.User, 0, len(dataModels))

	for _, dataModel := range dataModels {

		domainModels = append(domainModels, dataModel.ToDomainModel())
	}

	return domainModels, nil

}

func (repo *UserRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	dataModel := &model.User{}

	err := repo.DB.WithContext(ctx).Where("email = ?", email).First(dataModel).Error
	if err != nil {
		return nil, err
	}

	return dataModel.ToDomainModel(), nil

}

func (repo *UserRepo) Update(ctx context.Context, domainModel *domain.User) (*domain.User, error) {
	dataModel := &model.User{}

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

func (repo *UserRepo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	return repo.DB.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
}
