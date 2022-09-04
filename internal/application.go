package internal

import (
	"errors"
	"github.com/labstack/echo/v4"
	customErrors "golang_web_programming/errors"
	"net/http"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(ctx echo.Context) error {
	var createRequest CreateRequest

	if err := ctx.Bind(&createRequest); err != nil {
		str := customErrors.NewServiceError(err,customErrors.MessageParamsBinding, "400" ).Error()
		return errors.New(str)
	}

	if err := app.ValidateCreateRequest(createRequest); err != nil {
		return err
	}

	createResponse, err := app.repository.Create(&createRequest)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, createResponse)
}

func (app *Application) Update(ctx echo.Context) error {
	var request UpdateRequest

	//if err := ctx.Bind(&request); err != nil {
	//	return customErrors.ApiInternalServerError(customErrors.MessageParamsBinding)
	//}
	//
	//if request.ID == "" {
	//	return customErrors.ParamsValidationError(customErrors.MessageInvalidUserID)
	//}

	//if err := app.ValidateRequest(request); err != nil {
	//	return err
	//}

	updateResponse, err := app.repository.Update(&request)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, updateResponse)
}

func (app *Application) Delete(ctx echo.Context) error {
	id := ctx.Param("id")

	//if strings.Trim(id, " ") == "" {
	//	return customErrors.ParamsValidationError(customErrors.MessageInvalidUserID)
	//}

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

func (app *Application) ValidateCreateRequest(request CreateRequest) error {
	if request.UserName == "" {
		str := customErrors.NewServiceError(nil,customErrors.MessageInputUserName, "400" ).Error()

		return errors.New(str)
	}

	//if request.MembershipType == "" {
	//	return customErrors.ParamsValidationError(customErrors.MessageInputMembershipType)
	//}
	//if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
	//	return customErrors.ParamsValidationError(customErrors.MessageInvalidMembershipType)
	//}

	return nil
}
