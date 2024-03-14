package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/nazudis/disbursement/src/helper"
	"github.com/nazudis/disbursement/src/model/entity"
	"gorm.io/gorm"
)

type AuthSessionRepositoryImpl struct {
	db *gorm.DB
}

// Revoked implements AuthSessionRepository.
func (r *AuthSessionRepositoryImpl) Revoked(authSession *entity.AuthSession) error {
	authSession.IsRevoked = true
	return r.db.Save(authSession).Error
}

// Insert implements AuthSessionRepository.
func (r *AuthSessionRepositoryImpl) CreateAuthSession(ownerId uuid.UUID) (*entity.AuthSession, error) {
	now := time.Now()
	var authSession = &entity.AuthSession{
		OwnerID:   ownerId,
		IssuedAt:  now,
		ExpiredAt: now.Add(time.Minute * 5),
	}
	err := r.db.Create(&authSession).Error
	if err != nil {
		return nil, err
	}

	token, err := helper.GenerateJWTToken(map[string]interface{}{
		"client_id": ownerId.String(),
		"auth_id":   authSession.ID.String(),
		"exp":       authSession.ExpiredAt.Unix(),
		"iat":       authSession.IssuedAt.Unix(),
	})
	if err != nil {
		return nil, err
	}

	authSession.Token = token

	return authSession, nil
}

// FirstByID implements AuthSessionRepository.
func (r *AuthSessionRepositoryImpl) FirstByID(id string) (*entity.AuthSession, error) {
	var authSession entity.AuthSession
	if err := r.db.First(&authSession, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &authSession, nil
}

func NewAuthSessionRepository(db *gorm.DB) AuthSessionRepository {
	return &AuthSessionRepositoryImpl{db: db}
}
