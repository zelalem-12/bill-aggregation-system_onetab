package domain

import "time"

type Bill struct {
	Base
	linkedAccountID string
	amountDue       float64
	status          BillStatus
	dueDate         time.Time
	billedAt        time.Time
}

func (bill *Bill) SetLinkedAccountID(linkedAccountID string) {
	bill.linkedAccountID = linkedAccountID
}

func (bill *Bill) GetLinkedAccountID() string {
	return bill.linkedAccountID
}

func (bill *Bill) SetAmountDue(ammountDue float64) {
	bill.amountDue = ammountDue
}

func (bill *Bill) GetAmountDue() float64 {
	return bill.amountDue
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

func (bill *Bill) SetBilledAt(billedAt time.Time) {
	bill.billedAt = billedAt
}

func (bill *Bill) GetBilledAt() time.Time {
	return bill.billedAt
}

func NewBill(linkedAccountID string, amountDue float64, status BillStatus, dueDate, billedAt time.Time) *Bill {
	return &Bill{
		linkedAccountID: linkedAccountID,
		amountDue:       amountDue,
		status:          status,
		dueDate:         dueDate,
		billedAt:        billedAt,
	}
}
