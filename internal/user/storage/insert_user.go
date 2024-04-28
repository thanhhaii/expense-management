package userstorage

import (
	"context"

	usermodel "basicproject/internal/user/model"
)

func (storage *postgreStorage) CreateUser(
	ctx context.Context,
	data *usermodel.User,
) error {
	if err := storage.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}

	return nil
}
