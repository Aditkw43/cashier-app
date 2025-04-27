package service

import (
	"cashier-backend-go/internal/common/utils"
	"cashier-backend-go/internal/domain/auth/repository"
	"errors"
)

type AuthService interface {
	Login(username, password string) (string, error)
}

type authService struct {
	authRepo repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authService{authRepo: authRepo}
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.authRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
