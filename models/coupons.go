package models

import "time"

type Coupon struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Discount  int       `json:"discount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCoupon struct {
	Name     string `json:"name" validate:"required"`
	Discount int    `json:"discount" validate:"required"`
}
