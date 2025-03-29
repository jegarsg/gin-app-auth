package router

import (
	"GreatThanosApp/api/handler"
	"GreatThanosApp/api/middleware"
	"GreatThanosApp/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(
	db *gorm.DB, userHandler *handler.UserHandler, authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()
	authRepo := repository.NewAuthRepository(db)

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
	secure.Use(middleware.AuthMiddleware(authRepo))
	{
		secure.POST("/userbyemail/get", userHandler.GetUserByEmailHandler)
		secure.POST("/logout", authHandler.LogoutHandler)
	}

	return r
}
