package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func StatusUnauthorized(message string) error {
	return echo.NewHTTPError(http.StatusUnauthorized, message)
}

func StatusForbidden(message string) error {
	return echo.NewHTTPError(http.StatusForbidden, message)
}

func ApiRequestTooBigError(message string) error {
	return echo.NewHTTPError(http.StatusRequestEntityTooLarge, message)
}

func ApiInternalServerError(message string) error {
	return echo.NewHTTPError(http.StatusInternalServerError, message)
}

func ApiNotAcceptableError(message string) error {
	return echo.NewHTTPError(http.StatusNotAcceptable, message)
}

func ParamsValidationError(message string) error {
	return echo.NewHTTPError(http.StatusBadRequest, message)
}

func NoResultError(message string) error {
	return echo.NewHTTPError(http.StatusNotFound, message)
}

//func ApiParamValidError(err error) error {
//	return new(http.StatusBadRequest, 90001, err.Error())
//}