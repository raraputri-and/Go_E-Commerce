package customer

type Service interface{
	FindAll(UserID uint) ([]Customer, error)
	FindByID(ID int) (Customer, error)
	Create(ID uint, customerRequest CustomerRequest, UserID uint) (Customer, error)
	Update(ID int, customerRequest CustomerRequest) (Customer, error)
	Delete(ID int) (Customer, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) FindAll(UserID uint) ([]Customer, error) {
	customers, err := s.repository.FindAllByUser(UserID)
	return customers, err
}

func (s *service) FindByID(ID int) (Customer, error) {
	customer, err := s.repository.FindByID(ID)
	return customer, err
}

func (s *service) Create(ID uint, customerRequest CustomerRequest, UserID uint) (Customer, error) {
	customer := Customer {
		ID: UserID,
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

func (s *service) Delete(ID int) (Customer, error) {
	customer, err := s.repository.FindByID(ID)
	_, err = s.repository.Delete(customer)

	return customer, err
}