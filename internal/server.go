package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	log "github.com/sirupsen/logrus"
)

const _defaultPort = 8080

type Server struct {
	controller Controller
}

func NewDefaultServer() *Server {
	data := map[string]Membership{}
	service := NewService(*NewRepository(data))
	controller := NewController(*service)
	return &Server{
		controller: *controller,
	}
}

func (s *Server) Run() {
	e := echo.New()
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Printf("RequestHttpMethod=%s, RequestURI=%s, ResponseHttpStatusCode=%d", c.Request().Method,c.Request().RequestURI, c.Response().Status)
		log.Println("RequestBody="+ string(reqBody))
		log.Println("ResponseBody="+  string(resBody))
	}))


	e.GET("/", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	s.Routes(e)
	log.Fatal(e.Start(fmt.Sprintf(":%d", _defaultPort)))
}

func (s *Server) Routes(e *echo.Echo) {
	g := e.Group("/v1")
	RouteMemberships(g, s.controller)
}

func RouteMemberships(e *echo.Group, c Application) {
	e.POST("/memberships", c.Create)
	e.PUT("/memberships/:id", c.Update)
	e.GET("/memberships/:id", c.Get)
	e.DELETE("/memberships/:id", c.Delete)
	//e.POST("/memberships", c.Create, middleware.RequestIDWithConfig(middleware.RequestIDConfig{
	//	TargetHeader: "X-My-Request-Header",
	//}))
}
