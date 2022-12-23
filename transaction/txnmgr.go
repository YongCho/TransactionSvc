package transaction

import (
	"context"
	"fmt"

	"pismo.io/db"
	"pismo.io/db/sqlc"
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
func (t *TransactionMgr) CreateTransaction(accountID int32, opType OpType, amount float64) (Transaction, error) {
	switch opType {
	case Purchase:
		fallthrough
	case Withdrawal:
		if amount >= 0 {
			return Transaction{}, fmt.Errorf("amount must be negative for operation type %d", opType)
		}
	case CreditVoucher:
		if amount <= 0 {
			return Transaction{}, fmt.Errorf("amount must be positive for operation type %d", opType)
		}
	}

	// Money value is stored in cents to avoid float precision issue.
	amountCents := util.DollarToCents(amount)

	// Begin a new database transaction.
	conn := t.dbAdapter.GetDBHandle()
	tx, err := conn.Begin()
	if err != nil {
		return Transaction{}, err
	}
	defer tx.Rollback()
	qtx := t.dbAdapter.GetSqlcQueries().WithTx(tx)

	// Create the new transaction record.
	createParam := sqlc.CreateTransactionParams{
		AccountID:       accountID,
		OperationTypeID: int32(opType),
		Amount:          amountCents,
	}
	result, err := qtx.CreateTransaction(context.Background(), createParam)
	if err != nil {
		return Transaction{}, err
	}

	// Update the account balance.
	addBalanceParam := sqlc.AddAccountBalanceParams{
		ID:      accountID,
		Balance: amountCents,
	}
	_, err = qtx.AddAccountBalance(context.Background(), addBalanceParam)
	if err != nil {
		return Transaction{}, err
	}

	// Commit the database transaction.
	err = tx.Commit()
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
