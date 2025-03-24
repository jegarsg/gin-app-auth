package repository

import (
	"GreatThanosApp/models"

	"gorm.io/gorm"
)

// UserRepository handles database operations for users
type UserRepository struct {
	DB *gorm.DB // Inject GORM DB
}

// NewUserRepository initializes UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// ExistsByEmailOrPhone checks if a user exists by email or phone
func (r *UserRepository) ExistsByEmailOrPhone(email, phone string) bool {
	var count int64
	r.DB.Model(&models.User{}).Where("email = ? OR phone = ?", email, phone).Count(&count)
	return count > 0
}

// ExistsByUsername checks if a username already exists
func (r *UserRepository) ExistsByUsername(username string) bool {
	var count int64
	r.DB.Model(&models.User{}).Where("user_name = ?", username).Count(&count)
	return count > 0
}

// CreateUser saves a user in the database
func (r *UserRepository) CreateUser(user models.User) error {
	return r.DB.Create(&user).Error
}
