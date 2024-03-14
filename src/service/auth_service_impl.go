package service

import (
	"fmt"

	"github.com/nazudis/disbursement/src/model/entity"
	"github.com/nazudis/disbursement/src/repository"
)

type AuthServiceImpl struct {
	ClientRepository      repository.ClientRepository
	AuthSessionRepository repository.AuthSessionRepository
}

// GetAccessToken implements AuthService.
func (s *AuthServiceImpl) GetAccessToken(username string, password string) (*entity.AuthSession, error) {
	client, err := s.ClientRepository.FirstByID(username)
	if err != nil {
		return nil, err
	}
	if client == nil {
		return nil, fmt.Errorf("invalid client_id. client not found")
	}

	if !client.IsPasswordMatch(password) {
		return nil, fmt.Errorf("invalid client_secret. secret not match")
	}

	authSession, err := s.AuthSessionRepository.CreateAuthSession(client.ID)
	if err != nil {
		return nil, err
	}

	return authSession, nil
}

func NewAuthService(clietRepository repository.ClientRepository, authSessionRepository repository.AuthSessionRepository) AuthService {
	return &AuthServiceImpl{ClientRepository: clietRepository, AuthSessionRepository: authSessionRepository}
}
