INSERT INTO operation_type (id, description) VALUES
    (1, 'Normal Purchase'),
    (2, 'Purchase with Installments'),
    (3, 'Withdrawal'),
    (4, 'Credit Voucher')
ON CONFLICT (id) DO UPDATE SET description = excluded.description;
