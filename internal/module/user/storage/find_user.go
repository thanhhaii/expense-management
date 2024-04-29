package userstorage

import (
	usermodel "basicproject/internal/module/user/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (storage *postgreStorage) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*usermodel.User, error) {
	var (
		userRes usermodel.User
		db      = storage.db
	)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.WithContext(ctx).Where(conditions).First(&userRes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return &userRes, nil
}
