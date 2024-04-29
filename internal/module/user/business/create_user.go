package userbiz

import (
	"basicproject/internal/module/user/model"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
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

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return err
	}

	data.Password = string(hashPassword)
	return biz.store.CreateUser(ctx, data)
}
