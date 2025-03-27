package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func FormatPhone(phone string) (string, error) {
	if !regexp.MustCompile(`^[0-9]+$`).MatchString(phone) {
		return "", fmt.Errorf("invalid phone number: phone number must contain only digits")
	}

	if len(phone) < 10 || len(phone) > 12 {
		return "", fmt.Errorf("invalid phone number: phone number must be between 10 and 12 digits")
	}

	if strings.HasPrefix(phone, "0") {
		phone = "+62" + phone[1:]
	} else if !strings.HasPrefix(phone, "+62") {
		phone = "+62" + phone
	}

	return phone, nil
}
