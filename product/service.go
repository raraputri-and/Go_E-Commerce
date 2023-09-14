package product

type Service interface {
	// FindALL() ([]Product, error)
	FindAll(CustomerID uint) ([]Product, error)
	FindByID(ID int) (Product, error)
	// Create(productRequest ProductRequest) (Product, error)
	Create(productRequest ProductRequest, CustomerID uint) (Product, error)
	Update(ID int, productRequest ProductRequest) (Product, error)
	Delete(ID int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindALL(CustomerID uint) ([]Product, error) {
	products, err := s.repository.FindAllByCustomer(CustomerID)
	return products, err
}

func (s *service) FindByID(ID int) (Product, error) {
	product, err := s.repository.FindByID(ID)

	return product, err
}

func (s *service) Create(productRequest ProductRequest, CustomerID uint) (Product, error) {
	product := Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		CustomerID:  CustomerID}

	newProduct, err := s.repository.Create(product)
	return newProduct, err
}

func (s *service) Update(ID int, productRequest ProductRequest) (Product, error) {
	product, err := s.repository.FindByID(ID)

	if productRequest.Name != "" {
		product.Name = productRequest.Name
	}

	if productRequest.Description != "" {
		product.Description = productRequest.Description
	}

	if productRequest.Price != 0 {
		product.Price = productRequest.Price
	}

	updatedProduct, err := s.repository.Update(product)
	return updatedProduct, err
}

func (s *service) Delete(ID int) (Product, error) {
	product, err := s.repository.FindByID(ID)
	_, err = s.repository.Delete(product)

	return product, err
}
