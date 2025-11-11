package repository

import (
	"auth/config"
	"auth/internal/models"
)

type UserRepository struct{}

func (UserRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func (UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
