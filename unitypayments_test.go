package unitypayments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnityPayments(t *testing.T) {
	payments := New(nil, "", "", "")
	_, err := payments.ReceiveMobileMoney(ReceiveMobileMoneyRequest{
		CustomerMSISDN: "0240744225",
		Channel:        MTN,
		Amount:         .5,
		Description:    "test",
	})
	assert.NotNil(t, err)
	assert.Equal(t, "4010", err.(*ErrorResponse).ResponseCode)

	response, err := payments.ReceiveMobileMoney(ReceiveMobileMoneyRequest{
		CustomerName:       "Coolio",
		CustomerMSISDN:     "0240744225",
		Channel:            MTN,
		Amount:             .5,
		Description:        "test",
		PrimaryCallbackURL: "https://example.com",
	})
	assert.Nil(t, err)
	assert.Equal(t, "0001", response.ResponseCode)

	response, err = payments.SendMobileMoney(SendMobileMoneyRequest{
		RecipientName:      "Coolio",
		RecipientMSISDN:    "0240744225",
		Channel:            MTN,
		Amount:             .5,
		Description:        "test",
		PrimaryCallbackURL: "https://example.com",
	})
	assert.Nil(t, err)
	assert.Equal(t, "0001", response.ResponseCode)
}
