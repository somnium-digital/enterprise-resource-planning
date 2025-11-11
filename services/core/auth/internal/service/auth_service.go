package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"auth/internal/models"
	"auth/internal/repository"
	"auth/pkg/jwt"
)

type AuthService struct {
	Repo repository.UserRepository
}

func (s AuthService) Register(name, email, password string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{Name: name, Email: email, Password: string(hash)}
	return s.Repo.Create(&user)
}

func (s AuthService) Login(email, password string) (string, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid email or password")
	}

	return jwt.GenerateToken(user.ID, user.Email)
}
