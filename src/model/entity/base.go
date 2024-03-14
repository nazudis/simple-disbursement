package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseEntityUID struct {
	ID uuid.UUID `json:"id" gorm:"type:char(36);primarykey"`
}

func (m *BaseEntityUID) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
}
