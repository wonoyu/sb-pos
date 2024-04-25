-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS transaction_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
INSERT INTO transaction_type (name)
SELECT 'sale' WHERE NOT EXISTS (
    SELECT 'sale' FROM transaction_type
    WHERE name = 'sale'
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
INSERT INTO transaction_type (name)
SELECT 'refund' WHERE NOT EXISTS (
    SELECT 'refund' FROM transaction_type
    WHERE name = 'refund'
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    order_id INTEGER,
    customer_id INTEGER,
    ts_type_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_orders FOREIGN KEY(order_id) REFERENCES orders(id) ON DELETE CASCADE,
    CONSTRAINT fk_customer FOREIGN KEY(customer_id) REFERENCES customers(id) ON DELETE CASCADE,
    CONSTRAINT fk_transaction_type FOREIGN KEY(ts_type_id) REFERENCES transaction_type(id) ON DELETE RESTRICT
)
-- +migrate StatementEnd