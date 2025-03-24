package utils

import (
	"GreatThanosApp/config"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT generates a JWT token using email and username
func GenerateJWT(email string, username string) (string, error) {
	cfg := config.LoadConfig()

	// Create JWT claims including email, username, and expiration time
	claims := jwt.MapClaims{
		"email":    email,
		"username": username,
		"exp":      time.Now().Add(cfg.JWTExpireTime).Unix(),
	}

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateRefreshToken generates a JWT refresh token
func GenerateRefreshToken(email string, username string) (string, error) {
	cfg := config.LoadConfig()

	// Create JWT claims for the refresh token with a longer expiration time
	claims := jwt.MapClaims{
		"email":    email,
		"username": username,
		"exp":      time.Now().Add(cfg.JWTRefreshExpireTime).Unix(), // Longer expiration for refresh
	}

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*jwt.MapClaims, error) {
	cfg := config.LoadConfig()

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err) // Debug line
		return nil, err
	}

	if !token.Valid {
		fmt.Println("Invalid token") // Debug line
		return nil, errors.New("invalid token")
	}

	return &claims, nil
}
