package alarm_ex1

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

// 직접 구현한 mock 객체
type MockSMSClient1 struct {
}

func NewMockSMSClient1() *MockSMSClient1 {
	return &MockSMSClient1{}
}

func (m MockSMSClient1) Send(request SMSRequest) (SMSResponse, error) {
	return SMSResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}

// testify를 이용한 mock 객체
type MockSMSClient2 struct {
	mock.Mock
}

func NewMockSMSClient2() *MockSMSClient2 {
	return &MockSMSClient2{}

}

func (m MockSMSClient2) Send(request SMSRequest) (SMSResponse, error) {
	args := m.Called(request)
	return args.Get(0).(SMSResponse), args.Error(1)
}