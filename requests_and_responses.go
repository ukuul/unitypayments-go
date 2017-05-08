package unitypayments

import (
	"fmt"
)

type ReceiveMobileMoneyRequest struct {
	CustomerName         string
	CustomerMSISDN       string `json:"CustomerMsisdn"`
	CustomerEmail        string
	Channel              MobileMoneyProviders
	Amount               float64
	PrimaryCallbackURL   string `json:"PrimaryCallbackUrl"`
	SecondaryCallbackURL string `json:"SecondaryCallbackUrl"`
	Description          string
	ClientReference      string
	Token                string
}

type SendMobileMoneyRequest struct {
	RecipientName        string
	RecipientMSISDN      string `json:"RecipientMsisdn"`
	CustomerEmail        string
	Channel              MobileMoneyProviders
	Amount               float64
	PrimaryCallbackURL   string `json:"PrimaryCallbackUrl"`
	SecondaryCallbackURL string `json:"SecondaryCallbackUrl"`
	Description          string
	ClientReference      string
	Token                string
}

type ReceiveMobileMoneyData struct {
	TransactionID         string `json:"TransactionId"`
	ClientReference       string
	Description           string
	ExternalTransactionID string `json:"ExternalTransactionId"`
}

type MobileMoneyResponse struct {
	ResponseCode string
	Data         ReceiveMobileMoneyData
}

type ErrorResponse struct {
	ResponseCode string
	Message      string
	Errors       []ErrorData
}

type ErrorData struct {
	Field    string
	Messages []string
}

func (e ErrorResponse) Error() string {
	str := ""
	for _, item := range e.Errors {
		str += fmt.Sprintf("%+v", item)
	}
	return str
}
