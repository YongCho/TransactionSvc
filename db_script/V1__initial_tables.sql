-- This script contains all DDL statements to create the
-- database tables.

CREATE TABLE IF NOT EXISTS account (
    id SERIAL PRIMARY KEY,
    document_number TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS operation_type (
    id INT PRIMARY KEY,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS transaction (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL REFERENCES account(id),
    operation_type_id INT NOT NULL REFERENCES operation_type(id),
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
