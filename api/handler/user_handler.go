package handler

import (
	"GreatThanosApp/internal/service"
	"GreatThanosApp/models"
	"GreatThanosApp/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		pkg.SendResponse(c, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	createdUser, err := h.UserService.RegisterUser(user)
	if err != nil {
		pkg.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	pkg.SendResponse(c, http.StatusCreated, "User registered successfully!", createdUser)
}

func (h *UserHandler) GetUserByEmailHandler(c *gin.Context) {
	// Extract email from JWT claims stored in context
	email, exists := c.Get("Email")
	if !exists {
		pkg.SendResponse(c, http.StatusUnauthorized, "Unauthorized: No email found in token", nil)
		return
	}

	// Convert email to string
	emailStr, ok := email.(string)
	if !ok || emailStr == "" {
		pkg.SendResponse(c, http.StatusUnauthorized, "Invalid email in token", nil)
		return
	}

	// Fetch user by email
	response, err := h.UserService.GetUserByEmail(emailStr)
	if err != nil {
		pkg.SendResponse(c, http.StatusNotFound, err.Error(), nil)
		return
	}

	pkg.SendResponse(c, http.StatusOK, "User found", response)
}
