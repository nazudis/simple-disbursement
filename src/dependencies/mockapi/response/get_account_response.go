package response

type Account struct {
	ID            string `json:"id"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	BankShortCode string `json:"bank_short_code"`
	CreatedAt     string `json:"createdAt"`
}
