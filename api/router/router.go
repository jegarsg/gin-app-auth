package router

import (
	"GreatThanosApp/api/handler"
	"GreatThanosApp/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userHandler *handler.UserHandler, authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	// Public User routes
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/register", userHandler.Register)
	}

	// Authentication route
	r.POST("/api/loginsecure", authHandler.LoginSecureHandler)
	r.POST("/api/refreshtoken", authHandler.RefreshTokenHandler)

	// Secured routes (with JWT middleware)
	secure := r.Group("/api/secure")
	secure.Use(middleware.AuthMiddleware())
	{
		secure.POST("/userbyemail/get", userHandler.GetUserByEmailHandler)
	}

	return r
}
