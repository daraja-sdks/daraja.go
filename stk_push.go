package daraja

import (
	"encoding/json"
	"strconv"
)

type stkPush struct {
	_appRef            *mpesa
	_BusinessShortCode string
	_Password          string
	_Timestamp         string
	_TransactionType   string
	_Amount            string
	_PartyA            string
	_PartyB            string
	_PhoneNumber       string
	_CallBackURL       string
	_AccountReference  string
	_TransactionDesc   string
}

func stkPushBuilder(app *mpesa) *stkPush {
	return &stkPush{
		_appRef: app,
	}
}

func (s *stkPush) ShortCode(code string) *stkPush {
	s._BusinessShortCode = code
	s._PartyB = code
	return s
}

func (s *stkPush) PhoneNumber(no int) *stkPush {
	s._PhoneNumber = strconv.Itoa(no)
	s._PartyA = s._PhoneNumber
	return s
}

func (s *stkPush) Amount(amount int) *stkPush {
	s._Amount = strconv.Itoa(amount)
	return s
}

func (s *stkPush) AccountNumber(acc string) *stkPush {
	s._AccountReference = acc
	return s
}

func (s *stkPush) TransactionType(value transactionType) *stkPush {
	s._TransactionType = string(value)
	return s
}

func (s *stkPush) CallBackURL(url string) *stkPush {
	s._CallBackURL = url
	return s
}

func (s *stkPush) Description(value string) *stkPush {
	s._TransactionDesc = value
	return s
}

func (s *stkPush) Password(pass string) *stkPush {
	s._Password = pass
	return s
}

func (s *stkPush) Send() (*stkPushResponse, error) {
	body := s.marshall()
	token := s._appRef.GetAuthToken()
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	data, err := postRequest[stkPushResponse](s._appRef, routes["stkPush"], body, headers)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *stkPush) marshall() []byte {
	model := stkPushModel{
		BusinessShortCode: s._BusinessShortCode,
		Password:          s._Password,
		Timestamp:         s._Timestamp,
		TransactionType:   s._TransactionType,
		Amount:            s._Amount,
		PartyA:            s._PartyA,
		PartyB:            s._PartyB,
		PhoneNumber:       s._PhoneNumber,
		CallBackURL:       s._CallBackURL,
		AccountReference:  s._AccountReference,
		TransactionDesc:   s._TransactionDesc,
	}

	body, err := json.Marshal(model)
	if err != nil {
		panic("An internal error occurred. Please report this error at `https://github.com/ndaba1/daraja.go/issues`")
	}

	return body
}
