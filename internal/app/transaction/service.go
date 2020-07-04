package transaction

import (
	"fmt"

	"github.com/julioc98/pismo/internal/app/account"
)

type service struct {
	repo           Repository
	accountService account.Service
}

//NewService create new service factory
func NewService(r Repository, as account.Service) Service {
	return &service{
		repo:           r,
		accountService: as,
	}
}

// Create a Transaction
func (s service) Create(t *Transaction) (int, error) {
	if t.OperationID != 4 {
		t.Amount = t.Amount * -1
	}
	_, err := s.accountService.UpdateCreditLimit(t.AccountID, t.Amount)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(t)
}

// Get a Transaction
func (s service) Get(id int) (*Transaction, error) {
	return s.repo.Get(id)
}

func (s service) Check(t *Transaction) error {
	acc, err := s.accountService.Get(t.AccountID)
	if err != nil {
		return err
	}
	if t.Amount > acc.AvailableCreditLimit {
		return fmt.Errorf("credit not available for amount %f", t.Amount)
	}
	return nil
}
