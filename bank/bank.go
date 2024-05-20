package bank

import (
	"errors"
	"time"
)

type Account struct {
	UUID               string  `json:"id"`
	Balance            float64 `json:"balance"`
	TransactionHistory []Transaction
}

type Transaction struct {
	Balance           float64   `json:"balance"`
	TransactionAmount float64   `json:"transactionamount`
	Timestamp         time.Time `json:"timestamp"`
}

type BankHandler struct {
	Database map[string]Account
}

func (bh BankHandler) GetBankAccountByUUID(uuid string) (*Account, error) {
	for _, user := range bh.Database {
		if user.UUID == uuid {
			return &user, nil
		}
	}
	return nil, errors.New("No account found for " + uuid)
}

func (bh BankHandler) ApplyTransactionToAccount()
