package tokenfactory

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenFactory interface {
	CreateTokenWithClaims(email string) (string, error)
}

type tokenHelper struct {
	secretKey     []byte
	signingMethod jwt.SigningMethod
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateTokenFactory(
	secretKey []byte,
	signingMethod jwt.SigningMethod,
) *tokenHelper {
	return &tokenHelper{
		secretKey:     secretKey,
		signingMethod: signingMethod,
	}
}

func (tokenHelper *tokenHelper) CreateTokenWithClaims(email string) (string, error) {
	expirationTime := time.Now().Add(3 * 24 * time.Hour)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(tokenHelper.signingMethod, claims)
	tokenString, err := token.SignedString(tokenHelper.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
