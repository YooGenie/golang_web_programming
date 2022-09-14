package internal

import (
	"github.com/labstack/echo/v4"
	customErrors "golang_web_programming/errors"
)

func ValidateCreateRequest(request CreateRequest) error {
	if request.UserName == "" {
		return customErrors.ErrInputUserName
	}
	if request.MembershipType == "" {
		return customErrors.ErrInputMembershipType
	}
	if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		return customErrors.ErrInvalidMembershipType
	}

	return nil
}

func ValidateUpdateRequest(c echo.Context, request UpdateRequest) error {
	if request.ID == "" {
		return customErrors.ErrInvalidUserID
	}
	if request.UserName == "" {
		return customErrors.ErrInputUserName
	}
	if request.MembershipType == "" {
		return customErrors.ErrInputMembershipType
	}
	if !(request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco") {
		return customErrors.ErrInvalidMembershipType
	}

	return nil
}
