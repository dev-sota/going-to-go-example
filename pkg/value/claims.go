package value

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ispec-inc/going-to-go-example/pkg/config"
)

type Claims struct {
	UserID int64 `json:"uid"`
	jwt.StandardClaims
}

func NewClaims(userId int64) Claims {
	return Claims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.JWT.ExpireMin)).Unix(),
		},
	}
}
