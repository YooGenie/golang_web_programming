package user

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
	jwt.StandardClaims
}

func NewClaims(name string, isAdmin bool) Claims {
	return Claims{
		name,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour*24*365*10).Unix(),
		},
	}
}

func NewMemberClaims(name string) Claims {
	return NewClaims(name, false)
}

func NewAdminClaims(name string) Claims {
	return NewClaims(name, true)
}