-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    balance BIGINT,
    user_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_users FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS customer_orders (
    id SERIAL PRIMARY KEY,
    order_id INTEGER,
    customer_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_orders FOREIGN KEY(order_id) REFERENCES orders(id) ON DELETE CASCADE,
    CONSTRAINT fk_customers FOREIGN KEY(customer_id) REFERENCES customers(id) ON DELETE CASCADE
)
-- +migrate StatementEnd