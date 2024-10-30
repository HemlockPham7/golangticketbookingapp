package services

import (
	"context"

	"github.com/HemlockPham7/backend/models"
)

type AuthService struct {
	repository models.AuthRepository
}

func NewAuthService(repository models.AuthRepository) models.AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) Login(ctx context.Context, loginData *models.AuthCredentials) (string, *models.User, error) {

	return "", nil, nil
}

func (s *AuthService) Register(ctx context.Context, registerData *models.AuthCredentials) (string, *models.User, error) {

	return "", nil, nil
}
