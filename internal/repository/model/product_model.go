package model

import "time"

type ProductModel struct {
	Name        string
	Category    []string
	Quantity    float64
	Price       float64
	ImageUrl    string
	Sku         string
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
