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

type UserUseCase interface {
	RegisterUser(user models.User) (dto.RegisterUserResponse, error)
	GetUserByEmail(email string) (dto.GetUserByEmailResponse, error)
}

type userUseCase struct {
	userRepo repository.UserRepository // Use interface instead of pointer
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (uc *userUseCase) RegisterUser(user models.User) (dto.RegisterUserResponse, error) {
	formattedPhone, err := utils.FormatPhone(user.Phone)
	if err != nil {
		return dto.RegisterUserResponse{}, errors.New("invalid phone number format")
	}
	user.Phone = formattedPhone

	if uc.userRepo.ExistsByEmailOrPhone(user.Email, user.Phone) {
		return dto.RegisterUserResponse{}, errors.New("email or phone already registered")
	}

	user.UserName = uc.generateUniqueUsername(user.FullName)

	user.UserID = uuid.New()
	user.IsActive = false
	user.IsDeleted = false
	user.CreatedAt = time.Now()
	user.ModifiedAt = time.Now()
	user.CreatedBy = user.Email
	user.ModifiedBy = user.Email

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return dto.RegisterUserResponse{}, errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	if err := uc.userRepo.CreateUser(user); err != nil {
		return dto.RegisterUserResponse{}, errors.New("failed to save user")
	}

	response := dto.RegisterUserResponse{
		UserID:    user.UserID,
		UserName:  user.UserName,
		FullName:  user.FullName,
		Email:     user.Email,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
	}

	return response, nil
}

// Fixed function signature & removed infinite loop risk
func (uc *userUseCase) generateUniqueUsername(fullName string) string {
	baseUsername := strings.ReplaceAll(strings.ToLower(fullName), " ", "")
	username := baseUsername
	attempts := 0

	for uc.userRepo.ExistsByUsername(username) {
		attempts++
		username = utils.GenerateUsername(fullName)

		// Avoid infinite loop
		if attempts > 10 {
			username = username + uuid.New().String()[:8]
			break
		}
	}
	return username
}

func (uc *userUseCase) GetUserByEmail(email string) (dto.GetUserByEmailResponse, error) {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return dto.GetUserByEmailResponse{}, err // Return actual error
	}

	response := dto.GetUserByEmailResponse{
		UserID:    user.UserID,
		UserName:  user.UserName,
		FullName:  user.FullName,
		Email:     user.Email,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
	}
	return response, nil
}
