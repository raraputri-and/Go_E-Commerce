package product

type Service interface {
	// FindALL() ([]Product, error)
	FindAll(UserID uint) ([]Product, error)
	FindByID(ID int) (Product, error)
	// Create(productRequest ProductRequest) (Product, error)
	Create(productRequest ProductRequest, UserID uint) (Product, error)
	Update(ID int, productRequest ProductRequest) (Product, error)
	Delete(ID int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindALL(UserID uint) ([]Product, error) {
	products, err := s.repository.FindAllByUser(UserID)
	return products, err
}

func (s *service) FindByID(ID int) (Product, error) {
	product, err := s.repository.FindByID(ID)

	return product, err
}



