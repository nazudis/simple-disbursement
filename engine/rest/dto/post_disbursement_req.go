package dto

import "github.com/shopspring/decimal"

type PostDisbursementReq struct {
	ReferenceID           string          `json:"reference_id"`
	Description           string          `json:"description"`
	Amount                decimal.Decimal `json:"amount"`
	Type                  string          `json:"type"`
	BankShortCode         string          `json:"bank_short_code"`
	BankAccountNo         string          `json:"bank_account_no"`
	BankAccountHolderName string          `json:"bank_account_holder_name"`
}
