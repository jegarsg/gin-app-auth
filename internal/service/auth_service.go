package service

import (
	"GreatThanosApp/internal/dto"
	"GreatThanosApp/internal/usecase"

	"github.com/google/uuid"
)

type AuthService interface {
	Login(request dto.LoginRequest) (string, error)
	RefreshToken(refreshToken string) (string, string, error)
	Logout(userId uuid.UUID) error
}

type authService struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthService(authUseCase usecase.AuthUseCase) AuthService {
	return &authService{authUseCase: authUseCase}
}

func (s *authService) Login(request dto.LoginRequest) (string, error) {
	return s.authUseCase.Login(request.Email, request.Password)
}

func (s *authService) RefreshToken(refreshToken string) (string, string, error) {
	return s.authUseCase.RefreshNewTokens(refreshToken)
}

func (s *authService) Logout(userId uuid.UUID) error {
	return s.authUseCase.Logout(userId)
}
