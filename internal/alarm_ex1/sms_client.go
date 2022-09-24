package alarm_ex1

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type SMSClient interface {
	Send(request SMSRequest) (SMSResponse, error)
}

type SMSRequest struct {
	Title    string
	Body     string
	Receiver string
	Sender   string
}

type SMSResponse struct {
	Code    int
	Message string
}

// 실제 코드에서 사용할 client
type smsClient struct {
	client *http.Client
}

func (s smsClient) Send(request SMSRequest) (SMSResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return SMSResponse{}, err
	}

	httpResp, err := s.client.Post("sms_send.com", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return SMSResponse{}, err
	}

	var resp SMSResponse
	if err = json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return SMSResponse{}, err
	}

	return resp, nil
}
