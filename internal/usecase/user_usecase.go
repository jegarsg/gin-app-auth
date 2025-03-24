package usecase

import (
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

// RegisterUser handles user registration logic
func (uc *UserUseCase) RegisterUser(user models.User) (models.User, error) {

	formattedPhone, err := utils.FormatPhone(user.Phone)
	if err != nil {
		return models.User{}, errors.New("invalid phone number format")
	}
	user.Phone = formattedPhone

	if uc.UserRepo.ExistsByEmailOrPhone(user.Email, user.Phone) {
		return models.User{}, errors.New("email or phone already registered")
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
		return models.User{}, errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	if err := uc.UserRepo.CreateUser(user); err != nil {
		return models.User{}, errors.New("failed to save user")
	}

	return user, nil
}

func generateUniqueUsername(fullName string, uc *UserUseCase) string {
	baseUsername := strings.ReplaceAll(strings.ToLower(fullName), " ", "")
	username := baseUsername

	// Check if the username already exists and generate a unique one
	for uc.UserRepo.ExistsByUsername(username) {
		username = utils.GenerateUsername(fullName)
	}
	return username
}
