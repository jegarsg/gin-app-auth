package repository

import (
	"GreatThanosApp/config"
	"GreatThanosApp/models"
	"log"

	"github.com/google/uuid"
)

func GetAllUsers() []models.User {
	var users []models.User
	result := config.DB.Where("is_deleted = ?", false).Find(&users)
	if result.Error != nil {
		log.Fatalf("Error fetching users: %v", result.Error)
	}
	return users
}

func CreateUser(user models.User) {
	user.UserId = uuid.New() // Generate UUID for UserId
	result := config.DB.Create(&user)
	if result.Error != nil {
		log.Fatalf("Error creating user: %v", result.Error)
	}
}
