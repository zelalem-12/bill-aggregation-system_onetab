package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/adapter/persistence/postgres/model"
	repoPort "github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
	"gorm.io/gorm"
)

type BillRepo struct {
	DB *gorm.DB
}

func NewBillRepo(DB *gorm.DB) repoPort.BillRepo {
	return &BillRepo{
		DB: DB,
	}
}

func (repo *BillRepo) FindByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Bill, error) {
	var bills []model.Bill
	if err := repo.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&bills).Error; err != nil {
		return nil, err
	}

	var domainBills []*domain.Bill
	for _, bill := range bills {
		domainBills = append(domainBills, bill.ToDomainModel())
	}
	return domainBills, nil
}

func (repo *BillRepo) FindByProviderID(ctx context.Context, userID uuid.UUID, providerID uuid.UUID) ([]*domain.Bill, error) {
	var bills []model.Bill
	if err := repo.DB.WithContext(ctx).
		Where("user_id = ? AND provider_id = ?", userID, providerID).
		Find(&bills).Error; err != nil {
		return nil, err
	}

	var domainBills []*domain.Bill
	for _, bill := range bills {
		domainBills = append(domainBills, bill.ToDomainModel())
	}
	return domainBills, nil
}

func (repo *BillRepo) FindByProviderName(ctx context.Context, userID uuid.UUID, providerName string) ([]*domain.Bill, error) {
	var bills []model.Bill
	if err := repo.DB.WithContext(ctx).
		Where("user_id = ? AND provider_name = ?", userID, providerName).
		Find(&bills).Error; err != nil {
		return nil, err
	}

	var domainBills []*domain.Bill
	for _, bill := range bills {
		domainBills = append(domainBills, bill.ToDomainModel())
	}
	return domainBills, nil
}

func (repo *BillRepo) FindByID(ctx context.Context, billID uuid.UUID) (*domain.Bill, error) {
	var bill model.Bill
	if err := repo.DB.WithContext(ctx).Where("id = ?", billID).First(&bill).Error; err != nil {
		return nil, err
	}
	return bill.ToDomainModel(), nil
}

func (repo *BillRepo) Save(ctx context.Context, bill *domain.Bill) (*domain.Bill, error) {
	var dbBill model.Bill
	if err := dbBill.FromDomainModel(bill); err != nil {
		return nil, err
	}

	if err := repo.DB.WithContext(ctx).Save(&dbBill).Error; err != nil {
		return nil, err
	}
	return dbBill.ToDomainModel(), nil
}

func (repo *BillRepo) MarkAsPaid(ctx context.Context, billID uuid.UUID) error {
	return repo.DB.WithContext(ctx).
		Model(&model.Bill{}).
		Where("id = ?", billID).
		Update("status", "paid").Error
}

func (repo *BillRepo) Delete(ctx context.Context, billID uuid.UUID) error {
	return repo.DB.WithContext(ctx).
		Where("id = ?", billID).
		Delete(&model.Bill{}).Error
}

func (repo *BillRepo) Upsert(ctx context.Context, bills []*domain.Bill) error {

	for _, bill := range bills {
		var dbBill model.Bill
		if err := dbBill.FromDomainModel(bill); err != nil {
			return err
		}
		if err := repo.DB.WithContext(ctx).Save(&dbBill).Error; err != nil {
			return err
		}
	}
	return nil
}

func (repo *BillRepo) DeleteByProvider(ctx context.Context, userID uuid.UUID, providerID uuid.UUID) error {
	return repo.DB.WithContext(ctx).
		Where("user_id = ? AND provider_id = ?", userID, providerID).
		Delete(&model.Bill{}).Error
}
