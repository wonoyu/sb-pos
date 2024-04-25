package models

import "time"

type OrderStatus struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	StatusId   int    `json:"status_id"`
	StatusName string `json:"status_name"`
	Products   []OrderProductComplete
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateOrder struct {
	Products []OrderProducts `json:"products" validate:"required"`
}

type OrderProducts struct {
	Quantity  int `json:"quantity" validate:"required"`
	ProductId int `json:"product_id" validate:"required"`
}

type OrderProductComplete struct {
	ID              int       `json:"id"`
	ProductId       int       `json:"product_id"`
	ProductName     string    `json:"product_name"`
	ProductQuantity int       `json:"product_quantity"`
	ProductPrice    int64     `json:"product_price"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CompleteOrder struct {
	OrderId    int `json:"order_id"`
	CustomerId int `json:"customer_id"`
}
