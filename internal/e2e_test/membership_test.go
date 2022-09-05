package e2e_test

import (
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/labstack/echo/v4"
	"golang_web_programming/internal"
	"net/http"
	"testing"
)

func TestTossRecreate(t *testing.T) {
	echoServer := echo.New()
	internal.NewDefaultServer().Routes(echoServer)

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(echoServer),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Run("토스 멤버십을 신청한 후 삭제했다면, 다시 신청할 수 없다.", func(t *testing.T) {
		// given: 토스 멤버십을 신청한다.
		membershipCreateRequest := e.POST("/v1/memberships").
			WithJSON(internal.CreateRequest{
				UserName:       "andy",
				MembershipType: "toss",
			}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		// when: 토스 멤버십을 삭제한다.
		e.DELETE(fmt.Sprintf("/v1/memberships/%s", membershipCreateRequest.Value("ID").Raw())).
			Expect().
			Status(http.StatusOK)

		// then: 토스 멤버십을 다시 신청할 수 없다. 멤버십의 상태가 "탈퇴한 회원"이다.
		e.POST("/v1/memberships").
			WithJSON(internal.CreateRequest{
				UserName:       "andy",
				MembershipType: "toss",
			}).
			Expect().
			Status(http.StatusBadRequest).
			JSON().Object().
			Value("message").Equal("재가입할 수 없습니다.")
	})
}
