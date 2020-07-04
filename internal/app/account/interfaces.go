package account

// Repository Account interface
type Repository interface {
	Create(account *Account) (int, error)
	Get(id int) (*Account, error)
	Update(id int, account *Account) (*Account, error)
}

// Service Account interface
type Service interface {
	Create(account *Account) (int, error)
	Get(id int) (*Account, error)
	UpdateCreditLimit(id int, amount float64) (float64, error)
}
