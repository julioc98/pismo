package account

type service struct {
	repo Repository
}

//NewService create new service factory
func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

// Create a Account
func (s service) Create(a *Account) (int, error) {
	return s.repo.Create(a)
}

// Get a Account
func (s service) Get(id int) (*Account, error) {
	return s.repo.Get(id)
}
