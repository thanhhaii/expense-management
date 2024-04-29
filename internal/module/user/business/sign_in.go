package userbiz

import (
	usermodel "basicproject/internal/module/user/model"
	tokenfactory "basicproject/internal/token_factory"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type SignInStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

type signInBiz struct {
	store        SignInStore
	tokenService tokenfactory.TokenFactory
}

func CreateSignInBiz(store SignInStore, tokenService tokenfactory.TokenFactory) *signInBiz {
	return &signInBiz{
		store:        store,
		tokenService: tokenService,
	}
}

func (biz *signInBiz) SignIn(ctx context.Context, payload *usermodel.SignInModel) (string, error) {
	var (
		err   error
		token string
		user  *usermodel.User
	)
	user, err = biz.store.FindUser(ctx, map[string]interface{}{
		"email": payload.Email,
	})
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return "", err
	}

	token, err = biz.tokenService.CreateTokenWithClaims(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil

}
