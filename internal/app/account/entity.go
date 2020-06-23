package account

// Account entity
type Account struct {
	ID             int    `gorm:"primary_key" json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
