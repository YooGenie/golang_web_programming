package user

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct { //페이로드에 들어갈 값
	Name               string `json:"name"`
	IsAdmin            bool   `json:"is_admin"`
	jwt.StandardClaims        //기본적으로 사용되는 클레임값들이 있다.
}

func NewClaims(name string, isAdmin bool) Claims {
	return Claims{
		name,
		isAdmin,
		jwt.StandardClaims{
			//제한시간을 정해서 넣어준다.
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365 * 10).Unix(),
		},
	}
}

func NewMemberClaims(name string) Claims {
	return NewClaims(name, false)
}

func NewAdminClaims(name string) Claims {
	return NewClaims(name, true)
}
