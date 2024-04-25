package models

import "time"

type ProductCategory struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Price         int64     `json:"price"`
	StockQuantity int       `json:"stock_quantity"`
	CategoryId    int       `json:"category_id"`
	CategoryName  string    `json:"category_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateProduct struct {
	Name          string `json:"name" validate:"required"`
	Price         int64  `json:"price" validate:"required"`
	StockQuantity int    `json:"stock_quantity" validate:"required"`
	CategoryId    int    `json:"category_id" validate:"required"`
}
