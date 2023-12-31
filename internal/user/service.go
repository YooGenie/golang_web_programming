package user

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

var (
	DefaultSecret      = []byte("secret")
	ErrInvalidPassword = errors.New("invalid password")
)

type Service struct {
	secret []byte
}

func NewService(secret []byte) *Service {
	return &Service{secret: secret}
}

func (s Service) Login(name, password string) (LoginResponse, error) {
	if name != password {
		return LoginResponse{}, ErrInvalidPassword
	}

	claims := NewMemberClaims(name)
	if name == "admin" {
		claims = NewAdminClaims(name)
	}

	token, err := s.createToken(claims)
	if err != nil {
		return LoginResponse{}, err
	}
	return LoginResponse{Token: token}, nil
}

func (s Service) createToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 클레임: 키+벨류 => 페이로드
	return token.SignedString(s.secret) //시크릿값과 함께 토큰 값을 만들어 준다.
}
