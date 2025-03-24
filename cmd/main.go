package main

import (
	"GreatThanosApp/api/router"
	"GreatThanosApp/config"
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
	config.DB.AutoMigrate(&models.User{})

	// Start the server
	r := router.SetupRouter()

	// Swagger endpoint
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8090"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
