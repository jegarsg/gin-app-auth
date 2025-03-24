package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	JWTSecret            string
	JWTExpireTime        time.Duration
	JWTRefreshExpireTime time.Duration
}

func LoadConfig() *Config {
	jwtSecret := os.Getenv("JWT_SECRET")
	expireTimeStr := os.Getenv("JWT_EXPIRE_TIME")

	expireTime, err := strconv.Atoi(expireTimeStr)
	if err != nil || expireTime <= 0 {
		expireTime = 30
	}

	return &Config{
		JWTSecret:            jwtSecret,
		JWTExpireTime:        time.Duration(expireTime) * time.Second,
		JWTRefreshExpireTime: 24 * time.Hour,
	}
}
