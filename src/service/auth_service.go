package service

import "github.com/nazudis/disbursement/src/model/entity"

type AuthService interface {
	GetAccessToken(username, password string) (*entity.AuthSession, error)
}
