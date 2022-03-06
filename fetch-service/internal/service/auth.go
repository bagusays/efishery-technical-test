package service

import (
	"time"

	"github.com/bagusays/efishery-technical-test/internal"
	"github.com/bagusays/efishery-technical-test/internal/config"
	"github.com/bagusays/efishery-technical-test/internal/model"
	"github.com/golang-jwt/jwt"
)

//go:generate mockgen -destination=mock/mock_auth.go -package=mock github.com/bagusays/efishery-technical-test/internal/service Auth
type Auth interface {
	ValidateToken(tokenString string) (*model.UserClaim, error)
}

type auth struct {
	jwtSecret string
}

func NewAuth() Auth {
	return &auth{
		jwtSecret: config.GetConfig().JwtSecret,
	}
}

func (a auth) ValidateToken(tokenString string) (*model.UserClaim, error) {
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

type jwtClaim struct {
	model.User
	jwt.StandardClaims
}

func GenerateJWT(secret string, claimsData model.User) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &jwtClaim{
		User: claimsData,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
