package handler

import (
	"GreatThanosApp/config"
	"GreatThanosApp/models"
	"GreatThanosApp/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	fmt.Println("Received Token:", tokenString) // Debug line

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
		return
	}

	if len(tokenString) > 7 && strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
	}
	fmt.Println("Processed Token:", tokenString) // Debug line

	claims, err := utils.ValidateJWT(tokenString)
	if err != nil {
		fmt.Println("Token Validation Error:", err) // Debug line
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	email, ok := (*claims)["email"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload"})
		return
	}
	fmt.Println("Extracted Email from Token:", email) // Debug line

	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"fullName":    user.FullName,
		"username":    user.UserName,
		"email":       user.Email,
		"phone":       user.Phone,
		"isActive":    user.IsActive,
		"createdDate": user.CreatedDate,
	})
}
