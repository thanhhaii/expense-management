package transactionbiz

import (
	transactionmodel "basicproject/internal/module/transaction/model"
	"context"
	"fmt"
	"math/rand"
	"time"
)

type CreateTransactionStorage interface {
	CreateTransaction(ctx context.Context, data *transactionmodel.Transaction) error
}

type createTransactionBiz struct {
	storage CreateTransactionStorage
}

func NewCreateTransactionBiz(storage CreateTransactionStorage) *createTransactionBiz {
	return &createTransactionBiz{storage: storage}
}

func (biz *createTransactionBiz) CreateNewTransaction(
	ctx context.Context,
	data *transactionmodel.TransactionCreate,
) (string, error) {
	transactionCode := generateTransactionCode()

	dataCreate := &transactionmodel.Transaction{
		Sender:          data.Sender,
		Recipient:       data.Recipient,
		Amount:          data.Amount,
		Status:          0,
		Fee:             data.Fee,
		TransType:       data.TransType,
		Message:         data.Message,
		TransactionCode: transactionCode,
	}

	return transactionCode, biz.storage.CreateTransaction(ctx, dataCreate)
}

func generateTransactionCode() string {
	// Get current timestamp
	timestamp := time.Now().Unix()

	// Generate a random number (you can replace this with any unique identifier generation method)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := rand.Intn(1000) // Change 1000 to whatever range you want

	// Combine timestamp and random number to form the transaction code
	transactionCode := fmt.Sprintf("%d%d", timestamp, randomNum)

	return transactionCode
}
