package dto

type GetAccountVerificationRes struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankShortCode string `json:"bank_short_code"`
}
