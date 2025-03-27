package usecase

import (
	"GreatThanosApp/internal/dto"
	"GreatThanosApp/internal/repository"
	"GreatThanosApp/models"
	"GreatThanosApp/utils"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UserUseCase struct {
	UserRepo *repository.UserRepository
}

func NewUserUseCase(userRepo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{UserRepo: userRepo}
}

func (uc *UserUseCase) RegisterUser(user models.User) (dto.RegisterUserResponse, error) {

	formattedPhone, err := utils.FormatPhone(user.Phone)
	if err != nil {
		return dto.RegisterUserResponse{}, errors.New("invalid phone number format")
	}
	user.Phone = formattedPhone

	if uc.UserRepo.ExistsByEmailOrPhone(user.Email, user.Phone) {
		return dto.RegisterUserResponse{}, errors.New("email or phone already registered")
	}

	user.UserName = generateUniqueUsername(user.FullName, uc)

	user.UserId = uuid.New()
	user.IsActive = false
	user.IsDeleted = false
	user.CreatedDate = time.Now()
	user.ModifiedDate = time.Now()
	user.CreatedBy = user.Email
	user.ModifiedBy = user.Email

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return dto.RegisterUserResponse{}, errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	if err := uc.UserRepo.CreateUser(user); err != nil {
		return dto.RegisterUserResponse{}, errors.New("failed to save user")
	}

	response := dto.RegisterUserResponse{
		UserId:      user.UserId,
		UserName:    user.UserName,
		FullName:    user.FullName,
		Email:       user.Email,
		Phone:       user.Phone,
		IsActive:    user.IsActive,
		CreatedDate: user.CreatedDate,
	}

	return response, nil
}

func generateUniqueUsername(fullName string, uc *UserUseCase) string {
	baseUsername := strings.ReplaceAll(strings.ToLower(fullName), " ", "")
	username := baseUsername

	for uc.UserRepo.ExistsByUsername(username) {
		username = utils.GenerateUsername(fullName)
	}
	return username
}
