package initializers

import (
	"e-commerce/customer"
	"e-commerce/user"

	"gorm.io/gorm"
)

func SnycDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(user.User{}, customer.Customer{})
	return err
}
