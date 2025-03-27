package handler

import (
	"GreatThanosApp/internal/service"
	"GreatThanosApp/models"
	"GreatThanosApp/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
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
