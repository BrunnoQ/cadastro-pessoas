package validator

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	// E.164 phone number regex: +[country code][number]
	// Allows 1-3 digit country code and 4-14 digit phone number
	phoneRegex = regexp.MustCompile(`^\+[1-9]\d{0,2}\d{4,14}$`)
)

// ValidatePhone validates phone number in E.164 format
// Format: +[country code][number] (e.g., +5511999999999)
func ValidatePhone(phone string) error {
	if phone == "" {
		return fmt.Errorf("phone is required")
	}

	// Trim whitespace
	phone = strings.TrimSpace(phone)

	// Check format
	if !phoneRegex.MatchString(phone) {
		return fmt.Errorf("invalid phone format: must be in E.164 format (e.g., +5511999999999)")
	}

	// Check length (including + sign)
	if len(phone) < 8 || len(phone) > 16 {
		return fmt.Errorf("phone number must be between 8 and 16 characters")
	}

	return nil
}

// IsValidPhone checks if phone is valid (returns bool)
func IsValidPhone(phone string) bool {
	return ValidatePhone(phone) == nil
}

// FormatPhone ensures phone has + prefix
func FormatPhone(phone string) string {
	phone = strings.TrimSpace(phone)
	if !strings.HasPrefix(phone, "+") {
		return "+" + phone
	}
	return phone
}
