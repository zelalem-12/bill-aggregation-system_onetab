package domain

type BillStatus string

const (
	PAID   BillStatus = "paid"
	UNPAID BillStatus = "unpaid"
)

func (s BillStatus) String() string {
	return string(s)
}

func (s BillStatus) IsValid() bool {
	switch s {
	case PAID, UNPAID:
		return true
	}
	return false
}

func (s BillStatus) IsPaid() bool {
	return s == PAID
}

func (s BillStatus) IsUnpaid() bool {
	return s == UNPAID
}
