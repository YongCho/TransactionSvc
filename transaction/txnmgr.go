package transaction

import (
	"pismo.io/db"
	"pismo.io/util"
)

// TransactionMgr contains the main business logic for managing the
// accounts and the transactions.
type TransactionMgr struct {
	dbAdapter *db.DBAdapter
}

// NewTransactionMgr creates an instance of TransactionMgr.
func NewTransactionMgr(dbAdapter *db.DBAdapter) (*TransactionMgr, error) {
	return &TransactionMgr{dbAdapter: dbAdapter}, nil
}

// CreateAccount creates a new account.
func (t *TransactionMgr) CreateAccount(documentNumber string) (Account, error) {
	result, err := t.dbAdapter.CreateAccount(documentNumber)
	if err != nil {
		return Account{}, err
	}

	return Account{
		ID:             result.ID,
		DocumentNumber: result.DocumentNumber,
	}, nil
}

// GetAccount returns an existing account.
func (t *TransactionMgr) GetAccount(id int32) (Account, error) {
	result, err := t.dbAdapter.GetAccount(id)
	if err != nil {
		return Account{}, err
	}

	return Account{
		ID:             result.ID,
		DocumentNumber: result.DocumentNumber,
	}, nil
}

// CreateTransaction creates a new transaction record.
// Money value is stored in cents to avoid float precision issue.
func (t *TransactionMgr) CreateTransaction(accountID int32, opTypeID int32, amount float64) (Transaction, error) {
	amountCents := util.DollarToCents(amount)
	result, err := t.dbAdapter.CreateTransaction(accountID, opTypeID, amountCents)
	if err != nil {
		return Transaction{}, err
	}

	return Transaction{
		ID:              result.ID,
		AccountID:       result.AccountID,
		OperationTypeID: result.OperationTypeID,
		Amount:          result.Amount,
		CreatedAt:       result.CreatedAt,
	}, nil
}
