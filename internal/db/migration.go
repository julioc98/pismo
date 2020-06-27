package db

import (
	"github.com/jinzhu/gorm"
	"github.com/julioc98/pismo/internal/app/account"
	"github.com/julioc98/pismo/internal/app/transaction"
)

// Migrate migration BD
func Migrate(conn *gorm.DB) {
	// Migrate the schema
	conn.AutoMigrate(&account.Account{}, &transaction.Operation{}, &transaction.Transaction{})

	// Create an Account
	conn.Create(&account.Account{DocumentNumber: "540.100.249-92"})

	// Create Operactions
	conn.Create(&transaction.Operation{ID: 1, Description: "COMPRA A VISTA"})
	conn.Create(&transaction.Operation{ID: 2, Description: "COMPRA PARCELADA"})
	conn.Create(&transaction.Operation{ID: 3, Description: "SAQUE"})
	conn.Create(&transaction.Operation{ID: 4, Description: "PAGAMENTO"})

	//Create a Transaction
	conn.Create(&transaction.Transaction{
		AccountID:   1,
		OperationID: 1,
		Amount:      -50.55,
	})
}
