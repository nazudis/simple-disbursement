package dto

import "github.com/shopspring/decimal"

type PostDisbursementRes struct {
	ID                     string          `json:"id"`
	Type                   string          `json:"type"`
	Amount                 decimal.Decimal `json:"amount"`
	ReferenceID            string          `json:"reference_id"`
	RecipientAccountNumber string          `json:"recipient_account_number"`
	Status                 string          `json:"status"`
	Description            string          `json:"description"`
}
