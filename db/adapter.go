package db

import (
	"context"
	"database/sql"

	"pismo.io/db/sqlc"
)

// DBAdapter handles database interactions.
type DBAdapter struct {
	conn    *sql.DB
	queries *sqlc.Queries
}

// NewDBAdapter creates a new instance of DBAdapter.
func NewDBAdapter(conn *sql.DB) (*DBAdapter, error) {
	return &DBAdapter{
		conn:    conn,
		queries: sqlc.New(conn),
	}, nil
}

// ListAccounts returns all of the accounts in the DB.
// It can be extended with a pagination support in the future.
func (d *DBAdapter) ListAccounts() ([]sqlc.Account, error) {
	ctx := context.Background()
	return d.queries.ListAccounts(ctx)
}

// GetAccount returns a single account.
func (d *DBAdapter) GetAccount(accountID int32) (sqlc.Account, error) {
	ctx := context.Background()
	return d.queries.GetAccount(ctx, accountID)
}

// CreateAccount creates a new account.
func (d *DBAdapter) CreateAccount(documentNumber string) (sqlc.Account, error) {
	ctx := context.Background()
	return d.queries.CreateAccount(ctx, documentNumber)
}

// UpdateAccount updates an existing account.
func (d *DBAdapter) UpdateAccount(accountID int32, documentNumber string) (sqlc.Account, error) {
	ctx := context.Background()
	args := sqlc.UpdateAccountParams{
		ID:             accountID,
		DocumentNumber: documentNumber,
	}
	return d.queries.UpdateAccount(ctx, args)
}

// DeleteAccount deletes an existing account.
func (d *DBAdapter) DeleteAccount(accountID int32) error {
	ctx := context.Background()
	return d.queries.DeleteAccount(ctx, accountID)
}

// CreateTransaction creates a new transaction record.
func (d *DBAdapter) CreateTransaction(accountID int32, operationTypeID int32, amount int64) (sqlc.Transaction, error) {
	ctx := context.Background()
	args := sqlc.CreateTransactionParams{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
	}
	return d.queries.CreateTransaction(ctx, args)
}
