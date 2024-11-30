package dbmodels

const (
	Sending   AccountType = "sending"
	Receiving AccountType = "receiving"
)

type AccountType string

type Account struct {
	AccountNumber string      `json:"account_number"`
	AccountName   string      `json:"account_name"`
	IBAN          string      `json:"iban"`
	Address       string      `json:"address"`
	Amount        float64     `json:"amount"`
	Type          AccountType `json:"type"`
}
