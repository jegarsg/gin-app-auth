package router

import (
	"GreatThanosApp/api/handler"
	"GreatThanosApp/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public User routes
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/", handler.CreateUser)
		userGroup.POST("/register", handler.RegisterUser)
	}

	// Authentication route
	r.POST("/api/loginsecure", handler.LoginSecure)
	r.POST("/api/refreshtoken", handler.RefreshTokenHandler)

	// Secured routes (with JWT middleware)
	secure := r.Group("/api/secure")
	secure.Use(middleware.AuthMiddleware())
	{
		secure.GET("/users", handler.GetUsers)          // Example secure route
		secure.POST("/users/data", handler.GetUserData) // Secure data route
	}

	return r
}
