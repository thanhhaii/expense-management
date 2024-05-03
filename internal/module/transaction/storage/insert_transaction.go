package transactionstorage

import (
	transactionmodel "basicproject/internal/module/transaction/model"
	"context"
)

func (storage *postgreStorage) CreateTransaction(
	ctx context.Context,
	data *transactionmodel.Transaction,
) error {
	if err := storage.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}

	return nil
}
