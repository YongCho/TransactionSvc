-- This script contains the SQL queries that are used for
-- interacting with the database. It is used by sqlc framework
-- to generate Go code that interacts with the database.

-- name: ListAccounts :many
SELECT * FROM account
ORDER BY id;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1;

-- name: CreateAccount :one
INSERT INTO account (document_number)
VALUES ($1)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;

-- name: UpdateAccount :one
UPDATE account SET document_number = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE account SET balance = balance + $2
WHERE id = $1
RETURNING *;

-- name: CreateTransaction :one
INSERT INTO transaction (account_id, operation_type_id, amount)
VALUES ($1, $2, $3)
RETURNING *;
