package handler

import (
	"net/http"
	"time"

	"GreatThanosApp/config"
	"GreatThanosApp/models"
	"GreatThanosApp/utils"

	"github.com/gin-gonic/gin"
)

// loginsecure godoc
// @Summary Login Secure
// @Description Login Secure
// @Tags Login
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Users
// @Router /api/loginsecure [post]
func LoginSecure(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(user.Email, user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"expires": time.Now().Add(30 * time.Second).Unix(),
	})
}
