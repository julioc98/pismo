package account

import (
	gorm "github.com/jinzhu/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

//NewPostgresRepository create new postgres repository
func NewPostgresRepository(db *gorm.DB) Repository {
	return &postgresRepository{
		db,
	}
}

// Create Account
func (r *postgresRepository) Create(a *Account) (int, error) {
	if dbc := r.db.Create(a); dbc.Error != nil {
		return 0, dbc.Error
	}
	return a.ID, nil
}

// Get Account
func (r *postgresRepository) Get(id int) (*Account, error) {
	var account Account
	if dbc := r.db.First(&account, id); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &account, nil
}

func (r *postgresRepository) Update(id int, account *Account) (*Account, error) {
	if dbc := r.db.Save(account); dbc.Error != nil {
		return nil, dbc.Error
	}
	return account, nil
}
