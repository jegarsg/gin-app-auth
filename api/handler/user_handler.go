package handler

import (
	"GreatThanosApp/config"
	"GreatThanosApp/internal/service"
	"GreatThanosApp/models"
	"GreatThanosApp/pkg"
	"GreatThanosApp/utils"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetUsers godoc
// @Summary Get list of users
// @Description Get all users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /api/secure/users [get]
func GetUsers(c *gin.Context) {
	users := service.GetAllUsers()
	pkg.SendResponse(c, http.StatusOK, "Success", users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		pkg.SendResponse(c, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	service.CreateUser(user)
	pkg.SendResponse(c, http.StatusCreated, "User created", user)
}

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Format the phone number
	formattedPhone, err := utils.FormatPhone(user.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid phone number: %v", err)})
		return
	}
	user.Phone = formattedPhone

	// Check if email or phone already exists
	var existingUser models.User
	if err := config.DB.Where("email = ? OR phone = ?", user.Email, user.Phone).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email or phone already registered"})
		return
	}

	// Generate the base username from fullname
	baseUsername := strings.ReplaceAll(strings.ToLower(user.FullName), " ", "")
	user.UserName = baseUsername

	// Check if username already exists
	for {
		var duplicateUser models.User
		if err := config.DB.Where("user_name = ?", user.UserName).First(&duplicateUser).Error; err != nil {
			break
		}
		// Generate a new unique username if duplicate found
		user.UserName = utils.GenerateUsername(user.FullName)
	}

	// Generate a new UUID
	user.UserId = uuid.New()
	user.IsActive = false
	user.IsDeleted = false
	user.CreatedDate = time.Now()
	user.ModifiedDate = time.Now()
	user.CreatedBy = user.Email
	user.ModifiedBy = user.Email

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	// Save the user to the database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "User registered successfully!",
		"username": user.UserName,
		"phone":    user.Phone,
	})
}
