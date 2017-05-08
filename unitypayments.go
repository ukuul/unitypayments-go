package unitypayments

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"gopkg.in/jmcvetta/napping.v3"
)

// UnityPayments is the main payments object.
type UnityPayments struct {
	session *napping.Session
	baseURL string
}

// New creates a new instance of UnityPayments.
// Typically, we create one instance.
func New(client *http.Client, accountNumber, clientID, clientSecret string) *UnityPayments {
	b := []byte(fmt.Sprintf("%s:%s", clientID, clientSecret))
	token := base64.StdEncoding.EncodeToString(b)
	header := make(http.Header)
	header.Add("Accept", "application/json")
	header.Add("Authorization", "Basic "+token)
	header.Add("Cache-Control", "no-cache")
	header.Add("Content-Type", "application/json")

	if client == nil {
		client = http.DefaultClient
	}

	session := &napping.Session{
		Client: client,
		Header: &header,
	}

	return &UnityPayments{
		session: session,
		baseURL: "https://api.hubtel.com/v1/merchantaccount/merchants/" + accountNumber,
	}
}

// ReceiveMobileMoney initiates a request to receive payment via mobile money.
func (p UnityPayments) ReceiveMobileMoney(request ReceiveMobileMoneyRequest) (*MobileMoneyResponse, error) {
	response := new(MobileMoneyResponse)
	errResponse := new(ErrorResponse)
	res, err := p.session.Post(p.baseURL+"/receive/mobilemoney", &request, response, errResponse)
	if err != nil {
		return nil, err
	}
	if res.Status() >= 200 && res.Status() < 300 {
		return response, nil
	}
	return nil, errResponse
}

// SendMobileMoney send payment via mobile money.
func (p UnityPayments) SendMobileMoney(request SendMobileMoneyRequest) (*MobileMoneyResponse, error) {
	response := new(MobileMoneyResponse)
	errResponse := new(ErrorResponse)
	res, err := p.session.Post(p.baseURL+"/send/mobilemoney", &request, response, errResponse)
	if err != nil {
		return nil, err
	}
	if res.Status() >= 200 && res.Status() < 300 {
		return response, nil
	}
	return nil, errResponse
}
