package service

import (
	"GreatThanosApp/internal/dto"
	"GreatThanosApp/internal/usecase"
	"GreatThanosApp/models"
)

type UserService interface {
	RegisterUser(user models.User) (dto.RegisterUserResponse, error)
	GetUserByEmail(email string) (dto.GetUserByEmailResponse, error)
}

type userService struct {
	UserUseCase usecase.UserUseCase
}

func NewUserService(userUseCase usecase.UserUseCase) UserService {
	return &userService{UserUseCase: userUseCase}
}

func (s *userService) RegisterUser(user models.User) (dto.RegisterUserResponse, error) {
	return s.UserUseCase.RegisterUser(user)
}

func (s *userService) GetUserByEmail(email string) (dto.GetUserByEmailResponse, error) {
	return s.UserUseCase.GetUserByEmail(email)
}
