package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type JWTOptions struct {
	SigningKey string
	Expire     int
	Issuer     string
}

type JWTService struct {
	opt JWTOptions
}

func NewJWTService(opt JWTOptions) JWTService {
	return JWTService{opt: opt}
}

type UserClaims struct {
	model.ModelUser
	jwt.Claims
}

func (i *JWTService) GenerateToken(u *model.ModelUser) (string, error) {
	expire := time.Now().Add(time.Second * time.Duration(i.opt.Expire))
	claims := UserClaims{
		ModelUser: *u,
		Claims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
			Issuer:    i.opt.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(i.opt.SigningKey))

	return token, err
}

func (i *JWTService) ParseToken(token string) (u *model.ModelUser, err error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(i.opt.SigningKey), nil
	})
	c := claim.Claims.(jwt.MapClaims)
	toString, err := jsoniter.MarshalToString(c)
	if err != nil {
		return nil, err
	}
	u = &model.ModelUser{}
	err = jsoniter.Unmarshal([]byte(toString), u)
	if err != nil {
		return nil, err
	}
	return u, err
}
