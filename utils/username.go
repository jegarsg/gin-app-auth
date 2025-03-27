package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateUsername(fullname string) string {
	username := strings.ReplaceAll(strings.ToLower(fullname), " ", "")
	rand.Seed(time.Now().UnixNano())
	digits := rand.Intn(1000) // Three random digits
	letters := RandString(2)  // Two random letters
	return username + fmt.Sprintf("%02d%s", digits, letters)
}

func RandString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	rand.Seed(time.Now().UnixNano())
	result := make([]rune, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
