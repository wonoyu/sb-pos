package models

type Customer struct {
	ID       int    `json:"id"`
	Balance  int64  `json:"balance"`
	Username string `json:"user_name"`
	Email    string `json:"email"`
}

type TopupBalance struct {
	CustomerId int   `json:"customer_id"`
	Balance    int64 `json:"balance" validate:"required"`
}

type CustomerOrder struct {
	ID         int                    `json:"id"`
	CustomerId int                    `json:"customer_id"`
	OrderName  string                 `json:"order_name"`
	TotalPrice int64                  `json:"total_price"`
	Products   []OrderProductComplete `json:"products"`
}
