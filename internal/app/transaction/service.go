package transaction

type service struct {
	repo Repository
}

//NewService create new service factory
func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

// Create a Transaction
func (s service) Create(t *Transaction) (int, error) {
	if t.OperationID != 4 {
		t.Amount = t.Amount * -1
	}
	return s.repo.Create(t)
}

// Get a Account
func (s service) Get(id int) (*Transaction, error) {
	return s.repo.Get(id)
}
