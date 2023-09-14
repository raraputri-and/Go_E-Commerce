package initializers

import (
	"e-commerce/customer"
	"e-commerce/product"
	"e-commerce/user"

	"gorm.io/gorm"
)

func SnycDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(user.User{}, customer.Customer{}, product.Product{})
	return err
}
