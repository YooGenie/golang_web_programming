package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	customErrors "golang_web_programming/errors"
	"golang_web_programming/internal"
	"golang_web_programming/user"
)

type Middleware struct {
	membershipRepository internal.Repository
}

func NewMiddleware(membershipRepository internal.Repository) *Middleware {
	return &Middleware{membershipRepository: membershipRepository}
}

func JwtMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{Claims: &user.Claims{}, SigningKey: user.DefaultSecret})
}

func (m Middleware) ValidateAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userInfo := c.Get("user").(*jwt.Token)
		claims := userInfo.Claims.(*user.Claims)
		if !claims.IsAdmin {
			return customErrors.ErrUnauthorized
		}
		return next(c)
	}
}

func (m Middleware) ValidateMember(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userInfo := c.Get("user").(*jwt.Token)
		claims := userInfo.Claims.(*user.Claims)
		if claims.IsAdmin {
			return customErrors.ErrUnauthorized
		}

		id := c.Param("id")
		service := internal.Service{}
		membership, _ := service.GetByID(id)
		if claims.Name != membership.UserName {
			return customErrors.ErrUnauthorized
		}
		return next(c)
	}
}
