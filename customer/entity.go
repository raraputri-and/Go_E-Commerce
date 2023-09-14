package customer

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID   uint   `gorm:"ForeignKey:UserID"`
	Name string `gorm:"size:10"`
	// Product []product.Product
	UserID uint `gorm:"unique"`
}
