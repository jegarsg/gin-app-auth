package dto

import (
	"time"

	"github.com/google/uuid"
)

type RegisterUserResponse struct {
	UserId      uuid.UUID `json:"userId"`
	UserName    string    `json:"username"`
	FullName    string    `json:"fullname"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	IsActive    bool      `json:"isActive"`
	CreatedDate time.Time `json:"createdDate"`
}
