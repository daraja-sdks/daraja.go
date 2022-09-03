package daraja

type businessToCustomer struct {
	_appRef      *mpesa
	_Initiator   string
	_ShortCode   int32
	_PhoneNumber int32
	_CommandID   string
	_Remarks     string
	_TimeoutURL  string
	_ResultURL   string
	_Occassion   string
}

func b2cBuilder(app *mpesa) *businessToCustomer {
	return &businessToCustomer{
		_appRef:    app,
		_CommandID: "CustomerPayBillOnline",
	}
}

func (b *businessToCustomer) ShortCode(code int32) *businessToCustomer {
	b._ShortCode = code
	return b
}

func (b *businessToCustomer) PhoneNumber(value int32) *businessToCustomer {
	b._PhoneNumber = value
	return b
}

func (b *businessToCustomer) ResultURL(url string) *businessToCustomer {
	b._ResultURL = url
	return b
}
