package db

import (
	"github.com/jinzhu/gorm"
	"github.com/julioc98/pismo/internal/app/account"
)

// Migrate migration BD
func Migrate(conn *gorm.DB) {
	// Migrate the schema
	conn.AutoMigrate(&account.Account{})
}
