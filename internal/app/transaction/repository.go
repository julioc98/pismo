package transaction

import "github.com/jinzhu/gorm"

type postgresRepository struct {
	db *gorm.DB
}

//NewPostgresRepository create new postgres repository
func NewPostgresRepository(db *gorm.DB) Repository {
	return &postgresRepository{
		db,
	}
}

// Create Transaction
func (r *postgresRepository) Create(a *Transaction) (int, error) {
	if dbc := r.db.Create(a); dbc.Error != nil {
		return 0, dbc.Error
	}
	return a.ID, nil
}

// Get Transaction
func (r *postgresRepository) Get(id int) (*Transaction, error) {
	var transaction Transaction

	if dbc := r.db.Set("gorm:auto_preload", true).First(&transaction, id); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &transaction, nil
}
