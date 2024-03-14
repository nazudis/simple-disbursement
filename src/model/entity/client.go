package entity

import (
	"github.com/nazudis/disbursement/src/helper"
	"gorm.io/gorm"
)

type Client struct {
	BaseEntityUID
	Name   string `json:"name" gorm:"type:varchar(100);not null"`
	Secret string `json:"secret" gorm:"type:varchar(100);not null"`
	Timestamp
}

func (m *Client) IsPasswordMatch(password string) bool {
	return helper.CheckPasswordHash(password, m.Secret)
}

func (m *Client) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := helper.HashPassword(m.Secret)
	if err != nil {
		return
	}
	m.Secret = hash
	return
}

func (m *Client) AfterCreate(db *gorm.DB) error {
	wallet := &Wallet{OwnerID: m.ID}
	return db.Model(&Wallet{}).Create(wallet).Error
}
