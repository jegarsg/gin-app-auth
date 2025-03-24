package service

import (
	"GreatThanosApp/internal/usecase"
	"GreatThanosApp/models"
)

// UserService provides user-related services
type UserService struct {
	UserUseCase *usecase.UserUseCase // Inject UseCase Layer
}

// NewUserService initializes UserService
func NewUserService(userUseCase *usecase.UserUseCase) *UserService {
	return &UserService{UserUseCase: userUseCase}
}

// RegisterUser handles user registration
func (s *UserService) RegisterUser(user models.User) (models.User, error) {
	return s.UserUseCase.RegisterUser(user)
}
