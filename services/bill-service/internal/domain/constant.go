package domain

type BillStatus string

const (
	PAID    BillStatus = "paid"
	UNPAID  BillStatus = "unpaid"
	OVERDUE BillStatus = "overdue"
)
