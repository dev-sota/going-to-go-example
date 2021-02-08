package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/ispec-inc/going-to-go-example/pkg/config"
	"github.com/ispec-inc/going-to-go-example/pkg/presenter"
	"github.com/ispec-inc/going-to-go-example/pkg/registry"
	"github.com/ispec-inc/going-to-go-example/pkg/value"
	"github.com/ispec-inc/going-to-go-example/src/user"
)

type Auth struct {
	usecase user.Usecase
}

func NewAuth(repo registry.Repository) Auth {
	usecase := user.NewUsecase(repo)
	return Auth{usecase}
}

func (m Auth) VerifyToken(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tkn := r.Header.Get("Authorization")
		str, err := getTokenString(tkn)
		if err != nil {
			presenter.BadRequestError(w, err)
			return
		}

		claims, err := verifySignature(str)
		if err != nil {
			presenter.BadRequestError(w, err)
			return
		}

		inp := user.FindInput{
			ID: claims.UserID,
		}
		_, aerr := m.usecase.Find(inp)
		if aerr != nil {
			presenter.ApplicationException(w, aerr)
			return
		}

		ctx = context.WithValue(ctx, "uid", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func getTokenString(tkn string) (string, error) {
	if len(tkn) < 8 {
		return "", errors.New("Invalid Authorization header")
	}
	if tkn[:7] != "Bearer " {
		return "", errors.New("Authorization Header is not 'Bearer' token")
	}
	return tkn[7:], nil
}

func verifySignature(str string) (*value.Claims, error) {
	tkn, _ := jwt.ParseWithClaims(str, &value.Claims{}, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", tkn.Header["alg"])
		}
		return []byte(config.JWT.Secret), nil
	})

	claims, ok := tkn.Claims.(*value.Claims)
	if !(ok && tkn.Valid) {
		return &value.Claims{}, fmt.Errorf("unexpected signing method:")
	}

	return claims, nil
}
