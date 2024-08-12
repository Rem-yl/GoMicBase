// https://juejin.cn/post/7093035836689612836
package jwt_op

import (
	share "Account/Share"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var JWTKey = []byte("no rem no ley")

func GenJWTToken(name string) (string, error) {
	c := CustomClaims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "leyAccount",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(JWTKey)
}

func ParseJWTToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New(share.ErrInvalidToken)
}
