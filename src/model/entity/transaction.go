package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	TrxTypeBankTransfer = "bank_transfer"

	TrxStatusProcessing = "processing"
	TrxStatusCompleted  = "completed"
	TrxStatusFailed     = "failed"
)

type Transaction struct {
	BaseEntityUID
	OwnerID                uuid.UUID       `json:"owner_id" gorm:"type:varchar(36);not null"`
	Type                   string          `json:"type" gorm:"type:varchar(25);not null"`
	Amount                 decimal.Decimal `json:"amount" gorm:"type:decimal(64,15);default:0;not null"`
	ReferenceID            string          `json:"reference_id" gorm:"type:varchar(100);not null"`
	ExtReferenceID         string          `json:"ext_reference_id" gorm:"type:varchar(100);not null"`
	RecipientAccountNumber string          `json:"recipient_account_number" gorm:"type:varchar(100);not null"`
	Status                 string          `json:"status" gorm:"type:varchar(25);not null"`
	Description            string          `json:"description" gorm:"type:varchar(255)"`
	Timestamp
}
