-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS order_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
INSERT INTO order_status (name) 
SELECT 'paid' WHERE NOT EXISTS (
    SELECT 'paid' FROM order_status
    WHERE name = 'paid'
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
INSERT INTO order_status (name)
SELECT 'unpaid' WHERE NOT EXISTS (
    SELECT 'unpaid' FROM order_status
    WHERE name = 'unpaid'
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    status_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_order_status FOREIGN KEY(status_id) REFERENCES order_status(id) ON DELETE RESTRICT
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS order_products (
    id SERIAL PRIMARY KEY,
    order_id INTEGER,
    product_id INTEGER,
    product_quantity INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_orders FOREIGN KEY(order_id) REFERENCES orders(id) ON DELETE CASCADE,
    CONSTRAINT fk_products FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE RESTRICT
)
-- +migrate StatementEnd