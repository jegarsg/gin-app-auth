package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId       uuid.UUID `gorm:"type:uuid;primaryKey" json:"userId"`
	UserName     string    `gorm:"not null" json:"username"`
	FullName     string    `gorm:"not null" json:"fullname"`
	Password     string    `gorm:"not null" json:"password"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	Phone        string    `gorm:"uniqueIndex;not null" json:"phone"`
	IsActive     bool      `gorm:"default:false" json:"isActive"`
	CreatedDate  time.Time `gorm:"autoCreateTime" json:"createdDate"`
	CreatedBy    string    `gorm:"not null" json:"createdBy"`
	ModifiedDate time.Time `gorm:"autoUpdateTime" json:"modifiedDate"`
	ModifiedBy   string    `gorm:"not null" json:"modifiedBy"`
	IsDeleted    bool      `gorm:"default:false" json:"isDeleted"`
}
