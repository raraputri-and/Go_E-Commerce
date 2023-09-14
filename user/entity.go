package user

import (
	"e-commerce/customer"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Customer customer.Customer
}
