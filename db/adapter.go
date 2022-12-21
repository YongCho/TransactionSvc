package db

import (
	"context"
	"database/sql"

	"pismo.io/db/sqlc"
)

type DBAdapter struct {
	conn    *sql.DB
	queries *sqlc.Queries
}

func NewDBAdapter(conn *sql.DB) (*DBAdapter, error) {
	return &DBAdapter{
		conn:    conn,
		queries: sqlc.New(conn),
	}, nil
}

func (d *DBAdapter) ListAccounts() ([]sqlc.Account, error) {
	ctx := context.Background()
	return d.queries.ListAccounts(ctx)
}

func (d *DBAdapter) GetAccount(accountID int32) (sqlc.Account, error) {
	ctx := context.Background()
	return d.queries.GetAccount(ctx, accountID)
}

func (d *DBAdapter) CreateAccount(documentNumber string) (sqlc.Account, error) {
	ctx := context.Background()
	return d.queries.CreateAccount(ctx, documentNumber)
}

func (d *DBAdapter) UpdateAccount(accountID int32, documentNumber string) (sqlc.Account, error) {
	ctx := context.Background()
	args := sqlc.UpdateAccountParams{
		ID:             accountID,
		DocumentNumber: documentNumber,
	}
	return d.queries.UpdateAccount(ctx, args)
}

func (d *DBAdapter) DeleteAccount(accountID int32) error {
	ctx := context.Background()
	return d.queries.DeleteAccount(ctx, accountID)
}

func (d *DBAdapter) CreateTransaction(accountID int32, operationTypeID int32, amount int64) (sqlc.Transaction, error) {
	ctx := context.Background()
	args := sqlc.CreateTransactionParams{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
	}
	return d.queries.CreateTransaction(ctx, args)
}
