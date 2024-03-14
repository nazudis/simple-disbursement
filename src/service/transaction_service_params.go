package service

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type PostDisbursementParams struct {
	OwnerID               string
	ReferenceID           string
	Description           string
	Amount                decimal.Decimal
	Type                  string
	BankShortCode         string
	BankAccountNo         string
	BankAccountHolderName string
}

func (p *PostDisbursementParams) Validate() error {
	if p.OwnerID == "" {
		return fmt.Errorf("owner id is required")
	}
	if p.ReferenceID == "" {
		return fmt.Errorf("reference id is required")
	}
	if p.Amount.IsZero() {
		return fmt.Errorf("amount is required")
	}
	if p.Type == "" {
		return fmt.Errorf("type is required")
	}
	if p.BankShortCode == "" {
		return fmt.Errorf("bank short code is required")
	}
	if p.BankAccountNo == "" {
		return fmt.Errorf("bank account no is required")
	}
	if p.BankAccountHolderName == "" {
		return fmt.Errorf("bank account holder name is required")
	}
	return nil
}

type PostCallbackDisbursementParams struct {
	ID                    string
	Status                string
	CreatedAt             string
	ReferenceID           string
	Description           string
	Amount                decimal.Decimal
	Type                  string
	BankName              string
	BankShortCode         string
	BankAccountNo         string
	BankAccountHolderName string
}
