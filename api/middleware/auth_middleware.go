package middleware

import (
	"GreatThanosApp/internal/repository"
	"GreatThanosApp/utils"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(authRepo repository.AuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		if !authRepo.IsTokenValid(tokenString) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		username, err := getClaimValue(*claims, "username")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		email, _ := getClaimValue(*claims, "email") // Email is optional

		userId, _ := getClaimValue(*claims, "userId") // Email is optional

		c.Set("Username", username)
		c.Set("Email", email)
		c.Set("UserId", userId)
		c.Next()
	}
}

func getClaimValue(claims jwt.MapClaims, key string) (string, error) {
	value, exists := claims[key]
	if !exists {
		return "", fmt.Errorf("missing claim: %s", key)
	}

	strValue, ok := value.(string)
	if !ok {
		return "", errors.New("invalid claim format")
	}

	return strValue, nil
}
