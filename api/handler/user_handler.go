package handler

import (
	"GreatThanosApp/internal/service"
	"GreatThanosApp/models"
	"GreatThanosApp/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for user-related operations
type UserHandler struct {
	UserService *service.UserService // Inject UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// RegisterUser handles user registration
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		pkg.SendResponse(c, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	// Delegate business logic to the service layer
	createdUser, err := h.UserService.RegisterUser(user)
	if err != nil {
		pkg.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	pkg.SendResponse(c, http.StatusCreated, "User registered successfully!", createdUser)
}
