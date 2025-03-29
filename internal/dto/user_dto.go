package dto

import (
	"time"

	"github.com/google/uuid"
)

type RegisterUserResponse struct {
	UserID    uuid.UUID `json:"userId"`
	UserName  string    `json:"username"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type GetUserByEmailResponse struct {
	UserID    uuid.UUID `json:"userId"`
	UserName  string    `json:"username"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"CreatedAt"`
}
