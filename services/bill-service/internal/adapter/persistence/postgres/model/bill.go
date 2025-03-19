package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
)

type Bill struct {
	Base
	UserID          uuid.UUID `gorm:"type:uuid;not null"`
	LinkedAccountID string    `gorm:"type:uuid;not null"` // Each user has one account in each provider
	AmountDue       float64   `gorm:"not null"`
	Status          string    `gorm:"not null"`
	DueDate         time.Time `gorm:"not null"`
	BilledAt        time.Time `gorm:"not null"`
}

func (bill *Bill) FromDomainModel(domainBill *domain.Bill) error {
	if domainBill.GetID() != "" {
		billID, err := uuid.Parse(domainBill.GetID())
		if err != nil {
			return err
		}
		bill.ID = billID
	}
	bill.LinkedAccountID = domainBill.GetLinkedAccountID()
	bill.AmountDue = domainBill.GetAmountDue()
	bill.Status = string(domainBill.GetStatus())
	bill.DueDate = domainBill.GetDueDate()
	bill.BilledAt = domainBill.GetBilledAt()

	return nil
}

func (bill *Bill) ToDomainModel() *domain.Bill {
	domainBill := domain.Bill{}

	domainBill.SetID(bill.ID.String())
	domainBill.SetLinkedAccountID(bill.LinkedAccountID)
	domainBill.SetAmountDue(bill.AmountDue)
	domainBill.SetStatus(domain.BillStatus(bill.Status))
	domainBill.SetDueDate(bill.DueDate)
	domainBill.SetBilledAt(bill.BilledAt)

	domainBill.SetCreatedAt(bill.CreatedAt)
	domainBill.SetUpdatedAt(bill.UpdatedAt)
	if bill.DeletedAt.Valid {
		domainBill.SetDeletedAt(&bill.DeletedAt.Time)
	}
	return &domainBill
}
