package transaction

// Repository Account interface
type Repository interface {
	Create(account *Transaction) (int, error)
	Get(id int) (*Transaction, error)
}

// Service Account interface
type Service interface {
	Create(account *Transaction) (int, error)
	Get(id int) (*Transaction, error)
}
