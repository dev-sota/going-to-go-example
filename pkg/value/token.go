package value

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ispec-inc/going-to-go-example/pkg/apperror"
	"github.com/ispec-inc/going-to-go-example/pkg/config"
)

type Token string

func NewToken(userId int64) (Token, apperror.Error) {
	claims := jwt.MapClaims{}
	claims["uid"] = userId
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.JWT.ExpireMin)).Unix()

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tkn.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		return "", apperror.New(apperror.CodeError, err)
	}
	return Token(tokenString), nil
}
