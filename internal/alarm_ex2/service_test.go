package alarm_ex1

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestService(t *testing.T) {
	t.Run("전송 결과가 성공하면, retry하지 않고 에러를 리턴하지 않는다.", func(t *testing.T) {
		maxRetry := 3
		client := NewMockSMSClient()
		service := Service{client, maxRetry}

		receiver := "01022334444"
		client.On("Send", newSuccessSMSRequest(receiver)).
			Return(SMSResponse{http.StatusOK, http.StatusText(http.StatusOK)}, nil)

		err := service.Send(context.Background(), receiver)
		assert.Nil(t, err)
		client.AssertNumberOfCalls(t, "Send", 1)
	})

	t.Run("TooManyRequest 계속 발생하는 경우, 최대 retry count만큼 재시도한다.", func(t *testing.T) {
		maxRetry := 3
		client := NewMockSMSClient()
		service := Service{client, maxRetry}

		receiver := "01022334444"
		client.On("Send", newSuccessSMSRequest(receiver)).
			Return(SMSResponse{http.StatusTooManyRequests, "Too Many Request"}, nil)

		err := service.Send(context.Background(), receiver)
		assert.ErrorIs(t, err, ErrSMSFail)       // TODO
		client.AssertNumberOfCalls(t, "Send", 3) // TODO
	})

	t.Run("TooManyRequest 발생할 때마다 Retry하며, 도중에 성공할 경우 재시도하지 않는다.", func(t *testing.T) {
		/*
			client
			첫번째 요청의 응답: Code: http.StatusTooManyRequest, Message: http.StatusText(http.TooManyRequest)
			두번째 요청의 응답: Code: http.StatusOK, Message: http.StatusText(http.StatusOK)
			에러가 Nil이고, client를 2번 call했다는 것을 검증한다.
		*/

	})

	t.Run("client에서 internal server 에러가 발생한 경우, 재시도하지 않는다.", func(t *testing.T) {
		/*
			client
			응답: Code: http.InternalServerError, Message: http.StatusText(http.InternalServerError)
			에러가 SMSFailErr이고, client를 1번 call했다는 것을 검증한다.
		*/
	})
}

func TestServiceWithContext(t *testing.T) {
	t.Run("context가 취소되는 경우, 로직이 종료된다.", func(t *testing.T) {
		/*
			[테스트 로직]
			context.WithCancel 로 테스트
			client에 API이 call이 된 시점에 Cancel된다.
			client는 TooManyRequest를 리턴하고 있다.
			client가 재시도하려는 시점에 context가 Cancel되었다는 것을 알게 되고 에러를 리턴한다.
			service에서 리턴한 에러가 context에서 발생한 에러임을 검증한다.
		*/
	})

	t.Run("context 제한시간이 끝나는 경우, 로직이 종료된다.", func(t *testing.T) {
		/*
			[테스트 로직]
			context.WithTimeout : 500 ms
			client API Call하는데 1초가 소요된다.
			client는 TooManyRequest를 리턴하고 있다.
			client가 재시도하려는 시점에 context가 Timeout되었다는 것을 알게 되고 에러를 리턴한다.
			service에서 리턴한 에러가 context에서 발생한 에러임을 검증한다.
		*/
	})
}