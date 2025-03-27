package repository

import (
	"GreatThanosApp/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository initializes UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) ExistsByEmailOrPhone(email, phone string) bool {
	var count int64
	r.DB.Model(&models.User{}).Where("email = ? OR phone = ?", email, phone).Count(&count)
	return count > 0
}

func (r *UserRepository) ExistsByUsername(username string) bool {
	var count int64
	r.DB.Model(&models.User{}).Where("user_name = ?", username).Count(&count)
	return count > 0
}

func (r *UserRepository) CreateUser(user models.User) error {
	return r.DB.Create(&user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
