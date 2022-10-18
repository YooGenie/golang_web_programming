package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	customErrors "golang_web_programming/errors"
	"golang_web_programming/internal/member"
	"golang_web_programming/internal/user"
)

type Middleware struct {
	membershipRepository member.Repository
}

func NewMiddleware(membershipRepository member.Repository) *Middleware {
	return &Middleware{membershipRepository: membershipRepository}
}

func JwtMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{Claims: &user.Claims{}, SigningKey: user.DefaultSecret})
	//에코에서 제공하는 jwt 미들웨어가 있다.
	//미들웨어에서 토큰이 맞는지 아닌지 검증해준다.
	//제한 시간이 지나면 에코 JWT 미들웨어에서 만료시간이 지났다고 말한다
	//페이로드 변경했으면 잘못된 토큰이라고 알려준다.
}

func (m Middleware) ValidateAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userInfo := c.Get("user").(*jwt.Token)   //user라는 키에 jwt 토큰을 가져온다.
		claims := userInfo.Claims.(*user.Claims) //클레임을 보고 어드민인지 아닌지 체크 할 수 있다
		if !claims.IsAdmin {                     //어드민이 아니면 에러 처리
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
		service := member.Service{}
		membership, _ := service.GetByID(id)
		if claims.Name != membership.UserName {
			return customErrors.ErrUnauthorized
		}
		return next(c)
	}
}
