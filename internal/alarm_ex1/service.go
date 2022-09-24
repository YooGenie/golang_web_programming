package alarm_ex1

import (
	"errors"
	"net/http"
)

const _defaultSender = "0211112222"

var ErrSMSFail = errors.New("문자 전송에 실패했습니다")

type Service struct {
	smsClient SMSClient
}

func (service Service) Send(receiver string) error {
	resp, err := service.smsClient.Send(newSuccessSMSRequest(receiver))
	if err != nil {
		return err
	}
	if resp.Code == http.StatusOK {
		return nil
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