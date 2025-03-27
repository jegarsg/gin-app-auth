package usecase

import (
	"errors"
	"time"

	"GreatThanosApp/internal/dto"
	"GreatThanosApp/internal/repository"
	"GreatThanosApp/utils"
)

type AuthUseCase struct {
	UserRepo *repository.AuthRepository
}

func NewAuthUseCase(userRepo *repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{UserRepo: userRepo}
}

func (uc *AuthUseCase) Login(request dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := uc.UserRepo.FindByEmail(request.Email)
	if err != nil {
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(request.Password, user.Password) {
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.Email, user.UserName)
	if err != nil {
		return dto.LoginResponse{}, errors.New("failed to generate token")
	}

	return dto.LoginResponse{
		Message: "Login successful",
		Token:   token,
		Expires: time.Now().Add(30 * time.Second).Unix(),
	}, nil
}
