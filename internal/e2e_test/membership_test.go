package e2e_test

import (
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/labstack/echo/v4"
	"golang_web_programming/internal"
	"golang_web_programming/internal/member"
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
		membershipCreateRequest := e.POST("/v1/memberships"). // POST /v1/memberships
									WithJSON(member.CreateRequest{ //json 타입으로 보낸다. => Content-Type: application/json
				UserName:       "andy",
				MembershipType: "toss",
			}).                         // 바디값
			Expect().                   //여기 이후는 검증하는 것
			Status(http.StatusCreated). // 만약 베드리퀘스트로 오면 테스트는 실패
			JSON().Object()             //response 값을 json 포맷으로 바꿔준다.

		// when: 토스 멤버십을 삭제한다.
		token := e.POST("/v1/login").
			WithFormField("name", "admin").
			WithFormField("password", "admin").
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		e.DELETE(fmt.Sprintf("/v1/memberships/%s", membershipCreateRequest.Value("id").Raw())).
			//Value("id").Raw()는 스트럭스값이 아니라 json key값으로 가져온다.
			//.Row() => string으로 나온다.
			// DELETE /v1/memberships/:id 똑같이 보내진다
			WithHeader("authorization", fmt.Sprintf("bearer %s", token.Raw()["token"])).
			Expect().             //검증 부분
			Status(http.StatusOK) //OK말고 다른게 오면 테스트 코드 실패이다

		// then: 토스 멤버십을 다시 신청할 수 없다. 멤버십의 상태가 "탈퇴한 회원"이다.
		e.POST("/v1/memberships").
			WithJSON(member.CreateRequest{
				UserName:       "andy",
				MembershipType: "toss",
			}).
			Expect().
			Status(http.StatusInternalServerError).
			JSON().Object().
			Value("message").Equal("Internal Server Error")
		//JSON 키에 메시지 값에 "Internal Server Error" 적는다
	})

	t.Run("", func(t *testing.T) {
		// given: - 멤버십을 발급 받는다.
		membershipCreateRequest := e.POST("/v1/memberships").
			WithJSON(member.CreateRequest{
				UserName:       "jay",
				MembershipType: "toss",
			}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		// when: 	멤버십을 조회했을 때, 발급한 정보가 나온다
		token := e.POST("/v1/login").
			WithFormField("name", "jay").
			WithFormField("password", "jay").
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		id := membershipCreateRequest.Value("id").Raw()
		e.GET(fmt.Sprintf("/v1/memberships/%s", id)).
			WithHeader("authorization", fmt.Sprintf("bearer %s", token.Value("token").Raw())).
			Expect().
			Status(http.StatusOK).JSON().
			Equal(member.GetResponse{ //value 값을 비교한다.
				ID:             id.(string),
				UserName:       "jay",
				MembershipType: "toss",
			})
	})

	t.Run("멤버십의 주인만 멤버십을 조회할 수 있다", func(t *testing.T) {
		//Given: 멤버십을 생성한다
		membershipCreateRequest := e.POST("/v1/memberships").
			WithJSON(member.CreateRequest{
				UserName:       "kim",
				MembershipType: "toss",
			}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		id := membershipCreateRequest.Value("id").Raw()

		// When: 멤버십을 생성한 사용자가 로그인한다.
		token := e.POST("/v1/login").
			WithFormField("name", "kim").
			WithFormField("password", "kim").
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		// then: 사용자의 멤버십 단건 조회를 할 수 있다.
		e.GET(fmt.Sprintf("/v1/memberships/%s", id)).
			WithHeader("authorization", fmt.Sprintf("bearer %s", token.Value("token").Raw())).
			Expect().
			Status(http.StatusOK).
			JSON().Object().Equal(member.GetResponse{
			ID:             id.(string),
			UserName:       "kim",
			MembershipType: "toss",
		})
	})

	t.Run("Admin 사용자는 멤버십 전체 조회를 할 수 있다", func(t *testing.T) {
		//Given: 생성된 멤버십이 존재한다
		e.POST("/v1/memberships").
			WithJSON(member.CreateRequest{
				UserName:       "yona",
				MembershipType: "toss",
			}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		// When: Admin 사용자가 로그인한다
		token := e.POST("/v1/login").
			WithFormField("name", "admin").
			WithFormField("password", "admin").
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		// then: 멤버십 전체 조회를 할 수 있다.
		e.GET("/v1/memberships").
			WithHeader("authorization", fmt.Sprintf("bearer %s", token.Raw()["token"])).
			Expect().
			Status(http.StatusOK).
			JSON().Array()
	})

	t.Run("멤버십의 주인만 멤버십을 수정할 수 있다", func(t *testing.T) {
		//Given: 멤버십을 생성한다
		membershipCreateRequest := e.POST("/v1/memberships").
			WithJSON(member.CreateRequest{
				UserName:       "kim",
				MembershipType: "toss",
			}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		id := membershipCreateRequest.Value("id").Raw()

		// When: 멤버십을 생성한 사용자가 로그인한다.
		token := e.POST("/v1/login").
			WithFormField("name", "kim").
			WithFormField("password", "kim").
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		// then: 자기 멤버쉽을 수정할 수 있다.
		membershipUpdateRequest := e.PUT(fmt.Sprintf("/v1/memberships/%s", id)).
			WithJSON(member.UpdateRequest{
				ID:             id.(string),
				UserName:       "kim",
				MembershipType: "toss",
			}).
			WithHeader("authorization", fmt.Sprintf("bearer %s", token.Value("token").Raw())).
			Expect().
			Status(http.StatusCreated).
			JSON().Object().Equal(member.GetResponse{
			ID:             id.(string),
			UserName:       "kim",
			MembershipType: "toss",
		})
		fmt.Println("membershipUpdateRequest :", membershipUpdateRequest)
	})
}
