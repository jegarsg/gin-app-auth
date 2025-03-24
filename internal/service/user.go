package service

import (
	"GreatThanosApp/internal/repository"
	"GreatThanosApp/models"
)

func GetAllUsers() []models.User {
	return repository.GetAllUsers()
}

func CreateUser(user models.User) {
	repository.CreateUser(user)
}
