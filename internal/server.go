package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

const _defaultPort = 8080

type Server struct {
	controller Application
}

func NewDefaultServer() *Server {
	data := map[string]Membership{}
	controller := NewApplication(*NewRepository(data))
	return &Server{
		controller: *controller,
	}
}

func (s *Server) Run() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "RequestHttpMethod=${method}\nRequestURI=${uri}\nResponseHttpStatusCode=${status}\n",
	}))
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
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
