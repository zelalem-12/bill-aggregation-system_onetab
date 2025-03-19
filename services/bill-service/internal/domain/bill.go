package domain

import (
	"time"
)

type Bill struct {
	Base
	userID       string
	providerID   string
	providerName string
	amount       float64
	dueDate      time.Time
	status       BillStatus
}

func (bill *Bill) SetUserID(user_id string) {
	bill.userID = user_id
}

func (bill *Bill) GetUserID() string {
	return bill.userID
}

func (bill *Bill) SetProviderID(provider_id string) {
	bill.providerID = provider_id
}
func (bill *Bill) GetProviderID() string {
	return bill.providerID
}

func (bill *Bill) SetProviderName(provider_name string) {
	bill.providerName = provider_name
}

func (bill *Bill) GetProviderName() string {
	return bill.providerName
}

func (bill *Bill) SetAmount(ammount float64) {
	bill.amount = ammount
}

func (bill *Bill) GetAmount() float64 {
	return bill.amount
}

func (bill *Bill) SetStatus(status BillStatus) {
	bill.status = status
}

func (bill *Bill) GetStatus() BillStatus {
	return bill.status
}

func (bill *Bill) SetDueDate(dueDate time.Time) {
	bill.dueDate = dueDate
}

func (bill *Bill) GetDueDate() time.Time {
	return bill.dueDate
}

func NewBill(userId, providerId string, amount float64, status BillStatus, dueDate time.Time) *Bill {
	return &Bill{
		userID:     userId,
		providerID: providerId,
		amount:     amount,
		status:     status,
		dueDate:    dueDate,
	}
}
