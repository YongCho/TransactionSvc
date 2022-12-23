// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package sqlc

import (
	"context"
)

const addAccountBalance = `-- name: AddAccountBalance :one
UPDATE account SET balance = balance + $2
WHERE id = $1
RETURNING id, document_number, balance
`

type AddAccountBalanceParams struct {
	ID      int32
	Balance int64
}

func (q *Queries) AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, addAccountBalance, arg.ID, arg.Balance)
	var i Account
	err := row.Scan(&i.ID, &i.DocumentNumber, &i.Balance)
	return i, err
}

const createAccount = `-- name: CreateAccount :one
INSERT INTO account (document_number)
VALUES ($1)
RETURNING id, document_number, balance
`

func (q *Queries) CreateAccount(ctx context.Context, documentNumber string) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, documentNumber)
	var i Account
	err := row.Scan(&i.ID, &i.DocumentNumber, &i.Balance)
	return i, err
}

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transaction (account_id, operation_type_id, amount)
VALUES ($1, $2, $3)
RETURNING id, account_id, operation_type_id, amount, created_at
`

type CreateTransactionParams struct {
	AccountID       int32
	OperationTypeID int32
	Amount          int64
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction, arg.AccountID, arg.OperationTypeID, arg.Amount)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.OperationTypeID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, document_number, balance FROM account
WHERE id = $1
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(&i.ID, &i.DocumentNumber, &i.Balance)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many

SELECT id, document_number, balance FROM account
ORDER BY id
`

// This script contains the SQL queries that are used for
// interacting with the database. It is used by sqlc framework
// to generate Go code that interacts with the database.
func (q *Queries) ListAccounts(ctx context.Context) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(&i.ID, &i.DocumentNumber, &i.Balance); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE account SET document_number = $2
WHERE id = $1
RETURNING id, document_number, balance
`

type UpdateAccountParams struct {
	ID             int32
	DocumentNumber string
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount, arg.ID, arg.DocumentNumber)
	var i Account
	err := row.Scan(&i.ID, &i.DocumentNumber, &i.Balance)
	return i, err
}
