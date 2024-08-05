package jwt_op

import (
	"GoMicBase/0_Project/conf"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	TokenExpired     = errors.New("Token Expired")
	TokenNotValidYet = errors.New("Token Not Valid Yet")
	TokenMalformed   = errors.New("Token Malformed")
	TokenInvalid     = errors.New("Token Invalid")
)

type CustomClaims struct {
	jwt.StandardClaims
	Id          uint32
	Nickname    string
	AuthorityId int32
}

type JWT struct {
	SigninKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		SigninKey: []byte(conf.AppConf.JWTConfig.SingingKey),
	}
}

func (j *JWT) GenerateJWT(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(j.SigninKey)
	if err != nil {
		fmt.Printf("生成JWT错误: %s", err.Error())
		return "", err
	}

	return tokenStr, nil
}

func (j *JWT) ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigninKey, nil
	})
	if err != nil {
		if result, ok := err.(jwt.ValidationError); ok {
			if result.Errors != 0 {
				if jwt.ValidationErrorMalformed != 0 {
					return nil, TokenMalformed
				} else if jwt.ValidationErrorExpired != 0 {
					return nil, TokenExpired
				} else if jwt.ValidationErrorNotValidYet != 0 {
					return nil, TokenNotValidYet
				} else {
					return nil, TokenInvalid
				}
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}

		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}

}

func (j *JWT) RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.SigninKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(7 * 24 * time.Hour).Unix()
		return j.GenerateJWT(*&claims)
	}

	return "", TokenInvalid
}
