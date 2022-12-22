package transaction

import "time"

type OpType int32

const (
	Purchase      OpType = 1
	Installments  OpType = 2
	Withdrawal    OpType = 3
	CreditVoucher OpType = 4
)

type Account struct {
	ID             int32
	DocumentNumber string
}

type OperationType struct {
	ID          int32
	Description string
}

type Transaction struct {
	ID              int32
	AccountID       int32
	OperationTypeID int32
	Amount          int64
	CreatedAt       time.Time
}
