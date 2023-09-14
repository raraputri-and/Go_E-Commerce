package customer

import "gorm.io/gorm"

type Repository interface {
	FindAllByUser(UserID uint) (Customer, error)
	FindByID(ID int) (Customer, error)
	Create(customer Customer) (Customer, error)
	Update(customer Customer) (Customer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) FindAllByUser(UserID uint) (Customer, error) {
	var customer Customer
	err := r.db.Where("user_id = ?", UserID).Find(&customer).Error
	return customer, err
}

func (r *repository) FindByID(ID int) (Customer, error) {
	var customer Customer
	
	err := r.db.First(&customer, ID).Error

	return customer, err
}

func (r *repository) Create(customer Customer) (Customer, error) {
	err := r.db.Create(&customer).Error

	return customer, err
}

func (r *repository) Update(customer Customer) (Customer, error) {
	err := r.db.Save(&customer).Error

	return customer, err
}