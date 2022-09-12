package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang_web_programming/logo"
	"log"
	"net/http"
)

const _defaultPort = 8080

type Server struct {
	logoController       logo.Controller
	membershipController Controller
}

func NewDefaultServer() *Server {
	data := map[string]Membership{}
	membershipRepository := NewRepository(data)
	membershipService := NewService(*membershipRepository)
	membershipController := NewController(*membershipService)
	return &Server{
		membershipController: *membershipController,
		logoController:       *logo.NewController(),
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

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if err != nil {
			c.JSON(http.StatusBadRequest, "message : " + err.Error())
		}
		if echoErr, ok := err.(*echo.HTTPError); ok {
			if echoErr.Code == http.StatusInternalServerError {
				log.Print("Internal Server Error")

			}
		}

		e.DefaultHTTPErrorHandler(err, c)
	}

	e.GET("/", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	s.Routes(e)
	log.Fatal(e.Start(fmt.Sprintf(":%d", _defaultPort)))
}

func (s *Server) Routes(e *echo.Echo) {
	g := e.Group("/v1")
	RouteMemberships(g, s.membershipController)
	RouteLogo(g, s.logoController)
}

func RouteMemberships(e *echo.Group, c Controller) {
	e.POST("/memberships", c.Create)
	e.PUT("/memberships/:id", c.Update)
	e.GET("/memberships/:id", c.GetByID)
	e.DELETE("/memberships/:id", c.Delete)
	//e.POST("/memberships", c.Create, middleware.RequestIDWithConfig(middleware.RequestIDConfig{
	//	TargetHeader: "X-My-Request-Header",
	//}))
}


func RouteLogo(e *echo.Group, c logo.Controller) {
	e.GET("/logo", c.Get)
}