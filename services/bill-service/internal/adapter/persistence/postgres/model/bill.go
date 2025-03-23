package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/util"
)

type Bill struct {
	Base
	UserID       uuid.UUID `gorm:"type:uuid;not null;index"`
	ProviderID   uuid.UUID `gorm:"type:uuid;not null;index"`
	ProviderName string    `gorm:"not null;index"`
	Amount       float64   `gorm:"not null;default:0"`
	DueDate      time.Time `gorm:"not null"`
	Status       string    `gorm:"type:status;not null"`
	PaidDate     time.Time `gorm:"column:paid_at"`
}

func (bill *Bill) FromDomainModel(domainBill *domain.Bill) error {
	if domainBill.GetID() != "" {
		billID, err := util.ToUUID(domainBill.GetID())
		if err != nil {
			return err
		}
		bill.ID = billID
	}

	userId, err := util.ToUUID(domainBill.GetUserID())
	if err != nil {
		return err
	}
	bill.UserID = userId

	providerId, err := util.ToUUID(domainBill.GetProviderID())
	if err != nil {
		return err
	}
	bill.ProviderID = providerId

	bill.ProviderName = domainBill.GetProviderName()

	bill.Amount = domainBill.GetAmount()
	bill.Status = string(domainBill.GetStatus())
	bill.DueDate = domainBill.GetDueDate()
	bill.PaidDate = domainBill.GetPaidDate()

	return nil
}

func (bill *Bill) ToDomainModel() *domain.Bill {
	domainBill := domain.Bill{}

	domainBill.SetID(bill.ID.String())
	domainBill.SetUserID(util.ToString(bill.UserID))
	domainBill.SetProviderID(util.ToString(bill.ProviderID))
	domainBill.SetProviderName(bill.ProviderName)
	domainBill.SetAmount(bill.Amount)
	domainBill.SetStatus(domain.BillStatus(bill.Status))
	domainBill.SetDueDate(bill.DueDate)
	domainBill.SetPaidDate(bill.PaidDate)

	domainBill.SetCreatedAt(bill.CreatedAt)
	domainBill.SetUpdatedAt(bill.UpdatedAt)
	if bill.DeletedAt.Valid {
		domainBill.SetDeletedAt(&bill.DeletedAt.Time)
	}
	return &domainBill
}
