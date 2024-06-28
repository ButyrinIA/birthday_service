package service

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"rutube/internal/models"
	"rutube/internal/repository"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(user models.User) error {
	return s.repo.CreateUser(user)
}

var (
	ErrInvalidUsername = errors.New("неверное имя пользователя")
	ErrInvalidPassword = errors.New("неверный пароль")
)

func (s *AuthService) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvalidUsername
		}
		return nil, err
	}
	hash := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != bcrypt.ErrMismatchedHashAndPassword {
		return nil, ErrInvalidPassword
	}

	return user, nil
}
