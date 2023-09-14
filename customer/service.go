package customer

type Service interface{
	FindByID(ID int) (Customer, error)
	Create(customerRequest CustomerRequest, UserID uint) (Customer, error)
	Update(ID int, customerRequest CustomerRequest) (Customer, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) FindByID(ID int) (Customer, error) {
	customer, err := s.repository.FindByID(ID)
	return customer, err
}

func (s *service) Create(customerRequest CustomerRequest, UserID uint) (Customer, error) {
	customer := Customer {
		Name: customerRequest.Name,
		UserID: UserID,
	}

	newBook, err := s.repository.Create(customer)
	return newBook, err
}

func (s *service) Update(ID int, customerRequest CustomerRequest) (Customer, error) {
	customer, err := s.repository.FindByID(ID)

	if customerRequest.Name != "" {
		customer.Name = customerRequest.Name
	}

	newCustomer, err := s.repository.Update(customer)
	return newCustomer, err
}