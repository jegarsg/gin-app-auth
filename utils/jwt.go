package utils

import (
	"GreatThanosApp/config"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateJWT(userId uuid.UUID, email string, username string) (string, error) {
	cfg := config.LoadConfig()

	claims := jwt.MapClaims{
		"userId":   userId,
		"email":    email,
		"username": username,
		"exp":      time.Now().Add(cfg.JWTExpireTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(email string, username string) (string, error) {
	cfg := config.LoadConfig()

	claims := jwt.MapClaims{
		"email":    email,
		"username": username,
		"exp":      time.Now().Add(cfg.JWTRefreshExpireTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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
		fmt.Println("Error parsing token:", err)
		return nil, err
	}

	if !token.Valid {
		fmt.Println("Invalid token")
		return nil, errors.New("invalid token")
	}

	return &claims, nil
}
