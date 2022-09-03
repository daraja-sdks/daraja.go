package daraja

type accountBalance struct {
	_appRef *mpesa
}

func accountBalanceBuilder(app *mpesa) *accountBalance {
	return &accountBalance{
		_appRef: app,
	}
}
