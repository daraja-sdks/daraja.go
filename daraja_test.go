package daraja_test

import (
	"os"
	"testing"

	"github.com/ndaba1/daraja"
)

var (
	APP_KEY    = os.Getenv("APP_KEY")
	APP_SECRET = os.Getenv("APP_SECRET")
)

func Test_daraja_GetAuthToken(t *testing.T) {
	mpesa := daraja.Mpesa(APP_KEY, APP_SECRET, daraja.SANBOX)
	token := mpesa.GetAuthToken()
	if token == "" {
		t.Error("token is empty")
	}
}

func Test_daraja_STKPush(t *testing.T) {
	mpesa := daraja.Mpesa(APP_KEY, APP_SECRET, daraja.SANBOX)
	stkPush := mpesa.STKPush()
	res, e := stkPush.TransactionType(daraja.BuyGoods).
		CallBackURL("https://www.google.com").
		Description("Test").
		Password("password").
		ShortCode("600000").
		PhoneNumber(254708374149).
		Send()
	if e != nil {
		t.Error(e)
	}

	println(res.ResponseCode)
}
