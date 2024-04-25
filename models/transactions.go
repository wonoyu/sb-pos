package models

import "time"

type TransactionType struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTransaction struct {
	OrderId           int `json:"order_id"`
	CustomerId        int `json:"customer_id"`
	TransactionTypeId int `json:"transaction_type_id"`
}

type SalesTransaction struct {
	ID            int                    `json:"id"`
	OrderId       int                    `json:"order_id"`
	OrderName     string                 `json:"order_name"`
	CustomerId    int                    `json:"customer_id"`
	CustomerName  string                 `json:"customer_name"`
	Products      []OrderProductComplete `json:"products_sold"`
	TotalEarnings int64                  `json:"total_earnings"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

type ItemsSold struct {
	ProductName  string `json:"product_name"`
	SoldQuantity int    `json:"quantity"`
}

type TotalSales struct {
	ProductsSold int   `json:"products_sold"`
	Debit        int64 `json:"debit"`
}
