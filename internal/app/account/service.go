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

func (s service) UpdateCreditLimit(id int, amount float64) (float64, error) {
	acc, err := s.repo.Get(id)
	if err != nil {
		return 0, err
	}
	acc.AvailableCreditLimit = acc.AvailableCreditLimit + amount
	newAcc, err := s.repo.Update(id, acc)
	if err != nil {
		return 0, err
	}
	return newAcc.AvailableCreditLimit, nil
}
