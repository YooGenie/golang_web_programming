package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang_web_programming/internal"
	"golang_web_programming/logo"
	"golang_web_programming/middleware"
	"golang_web_programming/user"
	"log"
	"net/http"
)

const _defaultPort = 8080

type Server struct {
	logoController       logo.Controller
	membershipController internal.Controller
	userController       user.Controller
	userMiddleware       middleware.Middleware
}

func NewDefaultServer() *Server {
	data := map[string]internal.Membership{}
	membershipRepository := internal.NewRepository(data)
	membershipService := internal.NewService(*membershipRepository)
	membershipController := internal.NewController(*membershipService)
	return &Server{
		membershipController: *membershipController,
		logoController:       *logo.NewController(),
		userController:       *user.NewController(*user.NewService(user.DefaultSecret)),
		userMiddleware:       *middleware.NewMiddleware(*membershipRepository),
	}
}

func (s *Server) Run() {
	e := echo.New()
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Format: "RequestHttpMethod=${method}\nRequestURI=${uri}\nResponseHttpStatusCode=${status}\n",
	//}))
	//e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	//	log.Println("RequestBody="+ string(reqBody))
	//	log.Println("ResponseBody="+  string(resBody))
	//}))

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
	RouteUser(g, s.userController)
}

func RouteMemberships(e *echo.Group, c internal.Controller) {
	e.POST("/memberships", c.Create)
	e.PUT("/memberships/:id", c.Update,middleware.JwtMiddleware(), middleware.Middleware{}.ValidateMember)
	e.GET("/memberships/:id", c.GetByID, middleware.JwtMiddleware())
	e.GET("/memberships", c.GetList,middleware.JwtMiddleware(), middleware.Middleware{}.ValidateAdmin)
	e.DELETE("/memberships/:id", c.Delete,middleware.JwtMiddleware(), middleware.Middleware{}.ValidateAdmin)
	//e.POST("/memberships", c.Create, middleware.RequestIDWithConfig(middleware.RequestIDConfig{
	//	TargetHeader: "X-My-Request-Header",
	//}))
}


func RouteLogo(e *echo.Group, c logo.Controller) {
	e.GET("/logo", c.Get)
}

func RouteUser(e *echo.Group, c user.Controller) {
	e.POST("/login", c.Login)
}