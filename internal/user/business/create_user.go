package userbiz

import (
	"context"
	"errors"

	usermodel "basicproject/internal/user/model"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *usermodel.User) error
}

type createUserBiz struct {
	store CreateUserStorage
}

func NewCreateUserBiz(store CreateUserStorage) *createUserBiz {
	return &createUserBiz{
		store: store,
	}
}

func (biz *createUserBiz) CreateUser(ctx context.Context, data *usermodel.User) error {
	if data.Email == "" {
		return errors.New("Email cannot be blank")
	}

	if data.Password == "" {
		return errors.New("Passowrd cannot be blank")
	}

	return biz.store.CreateUser(ctx, data)
}
