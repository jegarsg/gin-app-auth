package router

import (
	"GreatThanosApp/api/handler"
	"GreatThanosApp/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	// Public User routes
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/register", userHandler.RegisterUser)
	}

	// Authentication route
	r.POST("/api/loginsecure", handler.LoginSecure)
	r.POST("/api/refreshtoken", handler.RefreshTokenHandler)

	// Secured routes (with JWT middleware)
	secure := r.Group("/api/secure")
	secure.Use(middleware.AuthMiddleware())
	{
		secure.POST("/users/data", handler.GetUserData) // Secure data route
	}

	return r
}
