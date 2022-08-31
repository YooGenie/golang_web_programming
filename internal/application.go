package internal

import (
	"github.com/labstack/echo/v4"
	customErrors "golang_web_programming/errors"
	"net/http"
	"strings"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(ctx echo.Context) error {
	var request Request

	if err := ctx.Bind(&request); err != nil {
		return customErrors.ApiInternalServerError(customErrors.MessageParamsBinding)
	}

	if err := app.ValidateRequest(request); err != nil {
		return err
	}

	createResponse, err := app.repository.Create(&request)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, createResponse)
}

func (app *Application) Update(ctx echo.Context) error {
	var request Request

	if err := ctx.Bind(&request); err != nil {
		return customErrors.ApiInternalServerError(customErrors.MessageParamsBinding)
	}

	if request.ID == "" {
		return customErrors.ParamsValidationError(customErrors.MessageInvalidUserID)
	}

	if err := app.ValidateRequest(request); err != nil {
		return err
	}

	updateResponse, err := app.repository.Update(&request)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, updateResponse)
}

func (app *Application) Delete(ctx echo.Context) error {
	id := ctx.Param("id")

	if strings.Trim(id, " ") == "" {
		return customErrors.ParamsValidationError(customErrors.MessageInvalidUserID)
	}

	if err := app.repository.Delete(id); err != nil {
		return err
	}

	return nil
}

func (app *Application) Get(ctx echo.Context) error {
	id := ctx.Param("id")

	getResponse, err := app.repository.GetOne(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, getResponse)
}

func (app *Application) ValidateRequest(request Request) error {
	if request.UserName == "" {
		return customErrors.ParamsValidationError(customErrors.MessageInputUserName)
	}
	if request.MembershipType == "" {
		return customErrors.ParamsValidationError(customErrors.MessageInputMembershipType)
	}
	if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		return customErrors.ParamsValidationError(customErrors.MessageInvalidMembershipType)
	}

	return nil
}
