package alarm_ex1

import (
	"context"
	"errors"
	"net/http"
)

const _defaultSender = "0211112222"

var (
	ErrSMSFail = errors.New("문자 전송에 실패했습니다")
)

type Service struct {
	smsClient     SMSClient
	maxRetryCount int
}

// 재시도하는 로직 사용
func (service Service) Send(ctx context.Context, receiver string) error {
	for i := 0; i < service.maxRetryCount; i++ { //maxRetryCount 만큼만 리트라이를 할 것이다.
		select {
		case <-ctx.Done(): //컨텍스트 취소 시그널을 알려주면
			return ctx.Err() //컨텍스트에서 내려 준 에러 응답하고 끝내준다
		default:
			resp, err := service.smsClient.Send(newSuccessSMSRequest(receiver)) //sms클라이언트에서 에러를 줬으면
			if err != nil {
				return err // dns (통신 X) 에러이면 다시 재시도 할 필요가 없다고 생각하고 바로 리턴을 해준다.
			}
			if resp.Code == http.StatusOK { //정상적으로 보내줬으면 재시도 할 필요가 없어서 리턴 nil를 해준다.
				return nil
			}
			if resp.Code == http.StatusTooManyRequests {  //TooManyRequests 인 경우에만 continue를 통해서 for문이 돈다.
				continue
			}
			return ErrSMSFail //internal Server Error, Bad request 재시도 안한다.
		}
	}
	return ErrSMSFail // maxRetryCount만큼 돌렸지만 성공도 다른 에러도 아닌 429만 나오는 경우니까 다시 시도 안하고 ErrSMSFail 에러를 내려준다.
}

func newSuccessSMSRequest(receiver string) SMSRequest {
	return SMSRequest{
		Title:    "가입 성공",
		Body:     "가입을 축하드립니다.",
		Receiver: receiver,
		Sender:   _defaultSender,
	}
}