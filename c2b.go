package daraja

type customerToBusiness struct {
	_appRef *mpesa
}

func c2bBuilder(app *mpesa) *customerToBusiness {
	return &customerToBusiness{
		_appRef: app,
	}
}
