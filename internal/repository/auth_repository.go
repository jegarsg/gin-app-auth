package repository

import (
	"GreatThanosApp/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

// NewUserRepository initializes UserRepository
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
