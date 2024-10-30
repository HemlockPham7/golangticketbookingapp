package repositories

import (
	"context"

	"github.com/HemlockPham7/backend/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) models.AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) RegisterUser(ctx context.Context, registerData *models.AuthCredentials) (*models.User, error) {

	return nil, nil
}

func (r *AuthRepository) GetUser(ctx context.Context, query interface{}, args ...interface{}) (*models.User, error) {

	return nil, nil
}
