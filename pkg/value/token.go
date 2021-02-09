package value

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ispec-inc/going-to-go-example/pkg/config"
)

type Token string

func NewToken(claims Claims) (Token, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tkn.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		return "", err
	}
	return Token(tokenString), nil
}
