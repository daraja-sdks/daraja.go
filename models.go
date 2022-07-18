package daraja

type transactionType string

const (
	PayBill  transactionType = "CustomerPayBillOnline"
	BuyGoods transactionType = "CustomerBuyGoodsOnline"
)

type authResponse struct {
	AccessToken string `json:"access_token"`
}

type stkPushModel struct {
	BusinessShortCode string
	Password          string
	Timestamp         string
	TransactionType   string
	Amount            string
	PartyA            string
	PartyB            string
	PhoneNumber       string
	CallBackURL       string
	AccountReference  string
	TransactionDesc   string
}

type stkPushResponse struct {
	MerchantRequestID string
	CheckoutRequestID string
	/**
	 * M-Pesa Result and Response Codes
	 *
	 * `0` - Success
	 * `1` - Insufficient Funds
	 * `2` - Less Than Minimum Transaction Value
	 * `3` - More Than Maximum Transaction Value
	 * `4` - Would Exceed Daily Transfer Limit
	 * `5` - Would Exceed Minimum Balance
	 * `6` - Unresolved Primary Party
	 * `7` - Unresolved Receiver Party
	 * `8` - Would Exceed Maxiumum Balance
	 * `11` - Debit Account Invalid
	 * `12` - Credit Account Invalid
	 * `13` - Unresolved Debit Account
	 * `14` - Unresolved Credit Account
	 * `15` - Duplicate Detected
	 * `17` - Internal Failure
	 * `20` - Unresolved Initiator
	 * `26` - Traffic blocking condition in place
	 *
	 */
	ResponseCode        string
	ResponseDescription string
	CustomerMessage     string
}
