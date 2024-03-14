package repository

import (
	"github.com/google/uuid"
	"github.com/nazudis/disbursement/src/model/entity"
)

type AuthSessionRepository interface {
	FirstByID(id string) (*entity.AuthSession, error)
	CreateAuthSession(ownerId uuid.UUID) (*entity.AuthSession, error)
	Revoked(authSession *entity.AuthSession) error
}
