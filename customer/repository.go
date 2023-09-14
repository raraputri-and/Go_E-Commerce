package customer

import "gorm.io/gorm"

type Repository interface {
	Create(customer Customer) (Customer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) Create(customer Customer) (Customer, error) {
	err := r.db.Create(&customer).Error

	return customer, err
}
