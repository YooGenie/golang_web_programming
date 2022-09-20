package internal

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	customErrors "golang_web_programming/errors"
	"golang_web_programming/user"
	"net/http"
	"strings"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) Create(c echo.Context) error {
	var request CreateRequest

	if err := c.Bind(&request); err != nil {
		return customErrors.ErrParamsBinding
	}

	if err := ValidateCreateRequest(request); err != nil {
		return err
	}

	createResponse, err := controller.service.Create(&request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, createResponse)
}

func (controller *Controller) Update(c echo.Context) error {
	var request UpdateRequest

	id := c.Param("id")

	if err := c.Bind(&request); err != nil {
		return customErrors.ErrParamsBinding
	}

	request.ID = id

	if err := ValidateUpdateRequest(request); err != nil {
		return err
	}

	updateResponse, err := controller.service.Update(&request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, updateResponse)
}

func (controller *Controller) GetByID(c echo.Context) error {
	id := c.Param("id")

	res, err := controller.service.GetByID(id)
	if err != nil {
		return echo.ErrInternalServerError
	}

	userInfo := c.Get("user").(*jwt.Token).Claims.(*user.Claims)
	if userInfo.Name != res.UserName && !userInfo.IsAdmin {
		return customErrors.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, res)
}

func (controller *Controller) Delete(c echo.Context) error {
	id := c.Param("id")

	if strings.Trim(id, " ") == "" {
		return customErrors.ErrInvalidUserID
	}

	if err := controller.service.Delete(id); err != nil {
		return err
	}

	return nil
}

func (controller *Controller) GetList(c echo.Context) error {

	res, err := controller.service.GetList()
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}
