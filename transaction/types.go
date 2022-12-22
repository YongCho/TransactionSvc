package transaction

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
