package validator

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	// RFC 5322 simplified email regex
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

// ValidateEmail validates email format according to RFC 5322
func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email is required")
	}

	// Trim whitespace
	email = strings.TrimSpace(email)

	// Check length
	if len(email) > 254 {
		return fmt.Errorf("email exceeds maximum length of 254 characters")
	}

	// Check format
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}

	// Split local and domain parts
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return fmt.Errorf("invalid email format")
	}

	local := parts[0]
	domain := parts[1]

	// Validate local part length
	if len(local) == 0 || len(local) > 64 {
		return fmt.Errorf("email local part must be between 1 and 64 characters")
	}

	// Validate domain part
	if len(domain) == 0 {
		return fmt.Errorf("email domain is required")
	}

	return nil
}

// IsValidEmail checks if email is valid (returns bool)
func IsValidEmail(email string) bool {
	return ValidateEmail(email) == nil
}
