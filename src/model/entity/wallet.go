package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	BaseEntityUID
	OwnerID uuid.UUID       `json:"owner_id" gorm:"type:varchar(36);not null"`
	Amount  decimal.Decimal `json:"amount" gorm:"type:decimal(64,15);not null"`
	Timestamp
}
