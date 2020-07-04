package transaction

import (
	"time"

	"github.com/julioc98/pismo/internal/app/account"
)

// Transaction entity
type Transaction struct {
	ID          int             `gorm:"primary_key" json:"id"`
	AccountID   int             `json:"account_id"`
	Account     account.Account `json:"account"`
	OperationID int             `json:"operation_type_id"`
	Operation   Operation       `json:"operation"`
	Amount      float64         `json:"amount"`
	Balance     float64         `json:"balance"`
	CreatedAt   time.Time       `json:"event_date"`
}

// Operation entity
type Operation struct {
	ID          int    `gorm:"primary_key"  json:"id"`
	Description string `json:"description"`
}
