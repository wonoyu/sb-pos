-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS product_category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
INSERT INTO product_category (name)
SELECT 'F&B' WHERE NOT EXISTS (
    SELECT 'F&B' FROM product_category
    WHERE name = 'F&B'
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    price BIGINT,
    stock_quantity INTEGER,
    category_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_product_category FOREIGN KEY(category_id) REFERENCES product_category(id) ON DELETE RESTRICT
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
INSERT INTO products (name, price, stock_quantity, category_id)
SELECT 'Susu SGM Explor', 1000000, 50, 1 WHERE NOT EXISTS (
    SELECT 'Susu SGM Explor', 1000000, 50, 1 FROM products
    WHERE name = 'Susu SGM Explor'
)
-- +migrate StatementEnd