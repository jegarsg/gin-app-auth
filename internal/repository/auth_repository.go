package repository

import (
	"GreatThanosApp/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	InvalidateToken(userId uuid.UUID) error
	SaveUserLogin(userId uuid.UUID, email, accessToken, refreshToken string, expiresAt time.Time) error
	IsTokenValid(token string) bool
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) InvalidateToken(userId uuid.UUID) error {
	return r.db.Where("user_id = ?", userId).Delete(&models.UserLogin{}).Error
}

func (r *authRepository) IsTokenValid(token string) bool {
	var count int64
	r.db.Model(&models.UserLogin{}).Where("access_token = ? AND expires_at > ?", token, time.Now()).Count(&count)
	return count > 0
}

func (r *authRepository) SaveUserLogin(userID uuid.UUID, email, accessToken, refreshToken string, expiresAt time.Time) error {
	var userLogin models.UserLogin

	err := r.db.Where("user_id = ?", userID).First(&userLogin).Error
	if err == nil {
		userLogin.AccessToken = accessToken
		userLogin.RefreshToken = refreshToken
		userLogin.ExpiresAt = expiresAt
		userLogin.ModifiedAt = time.Now()
		return r.db.Save(&userLogin).Error
	}

	userLogin = models.UserLogin{
		UserLoginID:  uuid.New(),
		UserID:       userID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
		CreatedAt:    time.Now(),
		CreatedBy:    email,
		ModifiedAt:   time.Now(),
		ModifiedBy:   email,
	}
	return r.db.Create(&userLogin).Error
}
