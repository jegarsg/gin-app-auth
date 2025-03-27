package service

import (
	"GreatThanosApp/internal/dto"
	"GreatThanosApp/internal/usecase"
)

type AuthService struct {
	AuthUseCase *usecase.AuthUseCase
}

func NewAuthService(authUseCase *usecase.AuthUseCase) *AuthService {
	return &AuthService{AuthUseCase: authUseCase}
}

func (s *AuthService) Login(request dto.LoginRequest) (dto.LoginResponse, error) {
	return s.AuthUseCase.Login(request)
}
