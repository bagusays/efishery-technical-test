package model

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleBasic Role = "BASIC"
)

type User struct {
	Phone      string    `json:"phone"`
	Name       string    `json:"name"`
	Role       Role      `json:"role"`
	UserName   string    `json:"userName"`
	Created_at time.Time `json:"created_at"`
}

type UserClaim struct {
	User
	jwt.StandardClaims
}
