package customer

type Service interface{
	Create(customerRequest CustomerRequest, UserID uint) (Customer, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) Create(customerRequest CustomerRequest, UserID uint) (Customer, error) {
	customer := Customer {
		Name: customerRequest.Name,
		UserID: UserID,
	}

	newBook, err := s.repository.Create(customer)
	return newBook, err
}