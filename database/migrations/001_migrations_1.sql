-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
INSERT INTO roles (name)
SELECT 'admin' WHERE NOT EXISTS (
    SELECT 'admin' FROM roles
    WHERE name = 'admin'
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
INSERT INTO roles (name)
SELECT 'customer' WHERE NOT EXISTS (
    SELECT 'customer' FROM roles
    WHERE name = 'customer'
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
INSERT INTO roles (name)
SELECT 'cashier' WHERE NOT EXISTS (
    SELECT 'cashier' FROM roles
    WHERE name = 'cashier'
)
-- +migrate StatementEnd

-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(256) UNIQUE,
    email VARCHAR(256) UNIQUE,
    password TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    role_id INTEGER,
    CONSTRAINT fk_roles FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE RESTRICT
)
-- +migrate StatementEnd