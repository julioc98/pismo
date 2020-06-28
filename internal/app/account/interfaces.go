package account

// Repository Account interface
type Repository interface {
	Create(account *Account) (int, error)
	Get(id int) (*Account, error)
}

// Service Account interface
type Service interface {
	Create(account *Account) (int, error)
	Get(id int) (*Account, error)
}
