package handler

import (
	"GreatThanosApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RefreshTokenHandler handles refreshing the JWT token
func RefreshTokenHandler(c *gin.Context) {
	// Get the refresh token from the request header
	refreshToken := c.GetHeader("Authorization")
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing refresh token"})
		return
	}

	// Validate the refresh token
	claims, err := utils.ValidateJWT(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Extract email and username from claims
	email, emailOk := (*claims)["email"].(string)
	username, usernameOk := (*claims)["username"].(string)
	if !emailOk || !usernameOk {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload"})
		return
	}

	// Generate a new access token
	newAccessToken, err := utils.GenerateJWT(email, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new access token"})
		return
	}

	// Generate a new refresh token
	newRefreshToken, err := utils.GenerateRefreshToken(email, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new refresh token"})
		return
	}

	// Send the new tokens
	c.JSON(http.StatusOK, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
