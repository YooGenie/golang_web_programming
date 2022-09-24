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

func (service Service) Send(ctx context.Context, receiver string) error {
	for i := 0; i < service.maxRetryCount; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			resp, err := service.smsClient.Send(newSuccessSMSRequest(receiver))
			if err != nil {
				return err
			}
			if resp.Code == http.StatusOK {
				return nil
			}
			if resp.Code == http.StatusTooManyRequests {
				continue
			}
			return ErrSMSFail
		}
	}
	return ErrSMSFail
}

func newSuccessSMSRequest(receiver string) SMSRequest {
	return SMSRequest{
		Title:    "가입 성공",
		Body:     "가입을 축하드립니다.",
		Receiver: receiver,
		Sender:   _defaultSender,
	}
}