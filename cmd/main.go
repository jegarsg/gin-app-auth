package main

import (
	"GreatThanosApp/api/handler"
	"GreatThanosApp/api/router"
	"GreatThanosApp/config"
	"GreatThanosApp/internal/repository"
	"GreatThanosApp/internal/service"
	"GreatThanosApp/internal/usecase"
	"GreatThanosApp/models"
	"log"

	// Swagger packages
	_ "GreatThanosApp/docs"
)

// @title Great Thanos App API
// @version 1.0
// @description This is a sample server for the Great Thanos App.
// @host localhost:8090
// @BasePath /api

func main() {
	// Connect to the database
	config.ConnectDB()

	// Automatically migrate the User model (create table if not exists)
	config.DB.AutoMigrate(&models.User{}, &models.UserLogin{})

	// Setup dependencies
	userRepo := repository.NewUserRepository(config.DB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userService := service.NewUserService(userUseCase)
	userHandler := handler.NewUserHandler(userService)

	authRepo := repository.NewAuthRepository(config.DB)
	authUseCase := usecase.NewAuthUseCase(authRepo)
	authService := service.NewAuthService(authUseCase)
	authHandler := handler.NewAuthHandler(authService) // Renamed variable

	// Start the server
	r := router.SetupRouter(config.DB, userHandler, authHandler) // Pass both handlers if needed

	// Swagger endpoint
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8090"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
