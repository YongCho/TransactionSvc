-- This script contains statements to seed the database
-- with the initial data.

INSERT INTO operation_type (id, description) VALUES
    (1, 'Normal Purchase'),
    (2, 'Purchase with Installments'),
    (3, 'Withdrawal'),
    (4, 'Credit Voucher')
ON CONFLICT (id) DO UPDATE SET description = excluded.description;
