package alarm_ex1

import (
	"github.com/stretchr/testify/mock"
)

// testify를 이용한 mock 객체
type MockSMSClient struct {
	mock.Mock
}

func NewMockSMSClient() *MockSMSClient {
	return &MockSMSClient{}

}

func (m *MockSMSClient) Send(request SMSRequest) (SMSResponse, error) {
	args := m.Called(request)
	return args.Get(0).(SMSResponse), args.Error(1)
}
