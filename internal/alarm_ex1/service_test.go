package alarm_ex1

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net"
	"net/http"
	"testing"
)

func TestService1(t *testing.T) {
	// 멤버십 생성이 성공했다는 문자를 보낸다.
	t.Run("문자 전송에 성공한다.", func(t *testing.T) {
		client := NewMockSMSClient1()
		service := Service{client}
		err := service.Send("01022334444")
		assert.Nil(t, err) // 메시지 전송에 성공했다면 에러가 nil이다 => 통과
	})
}

func TestService2(t *testing.T) {
	t.Run("문자 전송에 성공한다.", func(t *testing.T) {
		client := NewMockSMSClient2() // mock
		service := Service{client}

		receiver := "01022334444"
		client.On("Send", newSuccessSMSRequest(receiver)). //Send: 함수명
									Return(SMSResponse{http.StatusOK, "ok"}, nil)

		err := service.Send(receiver)
		assert.Nil(t, err)

	})

	t.Run("문자 전송에 성공한다.", func(t *testing.T) {
		client := NewMockSMSClient2() // mock
		service := Service{client}

		receiver := "01022334444"
		client.On("Send", mock.MatchedBy(func(req SMSRequest) bool { return req.Receiver == receiver })).
			Return(SMSResponse{http.StatusOK, "ok"}, nil)

		err := service.Send(receiver)
		assert.Nil(t, err)
	})

	t.Run("문자 전송에 성공한다. - 한번만", func(t *testing.T) {
		client := NewMockSMSClient2() // mock
		service := Service{client}

		receiver := "01022334444"
		client.On("Send", newSuccessSMSRequest(receiver)). //Send: 함수명
									Return(SMSResponse{http.StatusOK, "ok"}, nil).Once() //한번만이라서

		err := service.Send(receiver) //여기는 mock를 허용한다
		assert.Nil(t, err)

		err = service.Send(receiver) // 두번째는 어떤 값을 응답해야하는지 몰라서 에러가 발생한다
		assert.Nil(t, err)
	})

	t.Run("문자 전송에 실패한다. - 네트워크 통신 에러", func(t *testing.T) {
		expect := net.DNSError{Err: "dns error"}
		client := NewMockSMSClient2()
		service := Service{client}

		receiver := "01000000000"

		//mock를 친다.
		client.On("Send", newSuccessSMSRequest(receiver)).
			Return(SMSResponse{}, &expect)
		actual := service.Send(receiver)
		assert.ErrorIs(t, actual, &expect)
	})

	t.Run("문자 전송에 실패한다. - 전화번호가 유효하지 않을 때", func(t *testing.T) {
		/*
			Code: 400,
			Message: invalid phone number
		*/
		client := NewMockSMSClient2()
		service := Service{client}
		receiver := "010"
		client.On("Send", newSuccessSMSRequest(receiver)).
			Return(SMSResponse{http.StatusBadRequest, "invalid phone number"}, nil)
		err := service.Send(receiver)
		assert.ErrorIs(t, err, ErrSMSFail)
	})

}
