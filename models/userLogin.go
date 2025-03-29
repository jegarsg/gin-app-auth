package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserLogin struct {
	UserLoginID  uuid.UUID `gorm:"type:uuid;primaryKey" json:"userLoginId"`
	UserID       uuid.UUID `gorm:"type:uuid" json:"username"`
	AccessToken  string    `gorm:"unique"`
	RefreshToken string    `gorm:"unique"`
	ExpiresAt    time.Time `gorm:"index"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy    string    `gorm:"not null" json:"createdBy"`
	ModifiedAt   time.Time `gorm:"autoUpdateTime" json:"modifiedAt"`
	ModifiedBy   string    `gorm:"not null" json:"modifiedBy"`
}

func (u *UserLogin) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return nil
}
