package product

import "time"

type Product struct{
	ID int
	Name string
	Description string
	Price int
	CreatedAt time.Time
	UpdatedAt time.Time
	CustomerID uint
}