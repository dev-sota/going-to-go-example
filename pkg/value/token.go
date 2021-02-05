package value

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ispec-inc/going-to-go-example/pkg/apperror"
	"github.com/ispec-inc/going-to-go-example/pkg/config"
)

type Token string

func NewToken(user_id int64) (Token, apperror.Error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user_id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.JWT.ExpireMin)).Unix()

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := tkn.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		return "", apperror.New(apperror.CodeError, err)
	}
	return Token(signedString), nil
}
