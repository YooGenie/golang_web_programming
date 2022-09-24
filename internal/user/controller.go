package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller Controller) Login(c echo.Context) error {
	name := c.FormValue("name")
	password := c.FormValue("password")

	res, err := controller.service.Login(name, password)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
