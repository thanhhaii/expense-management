package transactionmodel

import "time"

type TransactionType int

const (
	TRANSFER TransactionType = iota + 1
	WITHDRAWAL
	DEPOSIT
	PAYMENT
	PURCHASE
)

type Transaction struct {
	Id        int `gorm:"column:id;"`
	Sender    int `gorm:"column:sender;"`
	Recipient int `gorm:"column:recipient;"`

	Amount          int             `gorm:"column:amount;"`
	Status          int             `gorm:"column:status;"`
	Fee             int             `gorm:"column:fee;"`
	TransType       TransactionType `gorm:"column:trans_type;"`
	Message         string          `gorm:"column:message;"`
	TransactionCode string          `gorm:"column:transaction_code;"`
	CreatedAt       *time.Time      `gorm:"column:created_at;"`
}

type TransactionCreate struct {
	Sender    int `json:"sender"`
	Recipient int `json:"recipient"`

	Amount    int             `json:"amount"`
	TransType TransactionType `json:"transType"`
	Message   string          `json:"message,omitempty"`
	Fee       int             `json:"fee,omitempty"`
}
