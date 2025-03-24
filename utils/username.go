package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// GenerateUsername generates a unique username based on the fullname
func GenerateUsername(fullname string) string {
	// Remove spaces and convert to lowercase
	username := strings.ReplaceAll(strings.ToLower(fullname), " ", "")

	// Generate a random suffix with 2 digits and 2 letters
	rand.Seed(time.Now().UnixNano())
	digits := rand.Intn(1000) // Two random digits
	letters := RandString(2)  // Two random letters
	return username + fmt.Sprintf("%02d%s", digits, letters)
}

// RandString generates a random string of a given length
func RandString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	rand.Seed(time.Now().UnixNano())
	result := make([]rune, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
