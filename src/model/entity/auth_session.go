package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthSession struct {
	BaseEntityUID
	OwnerID   uuid.UUID `json:"owner_id" gorm:"type:varchar(36);not null"`
	IssuedAt  time.Time `json:"issued_at" gorm:"type:timestamp;not null"`
	ExpiredAt time.Time `json:"expired_at" gorm:"type:timestamp;not null"`
	IsRevoked bool      `json:"is_revoked" gorm:"type:boolean;not null;default:false"`
	IsExpired bool      `json:"is_expired" gorm:"-"`
	Token     string    `json:"token" gorm:"-"`
	Timestamp
}

func (m *AuthSession) AfterFind(db *gorm.DB) (err error) {
	m.IsExpired = m.ExpiredAt.Before(time.Now())
	if m.IsExpired && !m.IsRevoked {
		m.IsRevoked = true
		db.Save(m)
	}
	return
}
