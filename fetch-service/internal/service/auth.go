package service

import (
	"github.com/bagusays/efishery-technical-test/internal"
	"github.com/bagusays/efishery-technical-test/internal/config"
	"github.com/bagusays/efishery-technical-test/internal/model"
	"github.com/golang-jwt/jwt"
)

type Auth struct {
	jwtSecret string
}

func NewAuth() *Auth {
	return &Auth{
		jwtSecret: config.GetConfig().JwtSecret,
	}
}

func (a Auth) ValidateToken(tokenString string) (*model.UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.jwtSecret), nil
	})

	if err != nil {
		return nil, internal.ErrInvalidToken
	}

	if claims, ok := token.Claims.(*model.UserClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, internal.ErrInvalidToken
}
