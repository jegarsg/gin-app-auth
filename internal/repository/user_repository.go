package repository

import (
	"GreatThanosApp/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) error
	FindByEmail(email string) (*models.User, error)
	ExistsByUsername(username string) bool
	ExistsByEmailOrPhone(email, phone string) bool
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) ExistsByEmailOrPhone(email, phone string) bool {
	var count int64
	r.DB.Model(&models.User{}).Where("email = ? OR phone = ?", email, phone).Count(&count)
	return count > 0
}

func (r *userRepository) ExistsByUsername(username string) bool {
	var count int64
	r.DB.Model(&models.User{}).Where("user_name = ?", username).Count(&count)
	return count > 0
}

func (r *userRepository) CreateUser(user models.User) error {
	return r.DB.Create(&user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
