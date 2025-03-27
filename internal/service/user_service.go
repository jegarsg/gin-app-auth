package service

import (
	"GreatThanosApp/internal/dto"
	"GreatThanosApp/internal/usecase"
	"GreatThanosApp/models"
)

type UserService struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserService(userUseCase *usecase.UserUseCase) *UserService {
	return &UserService{UserUseCase: userUseCase}
}

func (s *UserService) RegisterUser(user models.User) (dto.RegisterUserResponse, error) {
	return s.UserUseCase.RegisterUser(user)
}
