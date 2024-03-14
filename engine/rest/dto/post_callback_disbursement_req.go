package dto

import "github.com/shopspring/decimal"

type PostCallbackDisbursementReq struct {
	ID                    string          `json:"id"`
	Status                string          `json:"status"`
	CreatedAt             string          `json:"created_at"`
	ReferenceID           string          `json:"reference_id"`
	Description           string          `json:"description"`
	Amount                decimal.Decimal `json:"amount"`
	Type                  string          `json:"type"`
	BankName              string          `json:"bank_name"`
	BankShortCode         string          `json:"bank_short_code"`
	BankAccountNo         string          `json:"bank_account_no"`
	BankAccountHolderName string          `json:"bank_account_holder_name"`
}
