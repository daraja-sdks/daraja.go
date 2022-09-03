package daraja

type businessToCustomer struct {
	_appRef *mpesa
}

func b2cBuilder(app *mpesa) *businessToCustomer {
	return &businessToCustomer{
		_appRef: app,
	}
}
