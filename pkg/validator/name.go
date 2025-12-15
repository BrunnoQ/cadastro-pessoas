package validator

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var (
	// Name regex: Unicode letters, spaces, hyphens, apostrophes
	nameRegex = regexp.MustCompile(`^[\p{L}\s\-']+$`)
)

const (
	// MinNameLength is the minimum length for names
	MinNameLength = 2
	// MaxNameLength is the maximum length for names
	MaxNameLength = 100
)

// ValidateName validates person name or surname
// Allows: Unicode letters, spaces, hyphens, apostrophes
// Rejects: Numbers, special characters (except hyphen and apostrophe)
func ValidateName(name string, fieldName string) error {
	if name == "" {
		return fmt.Errorf("%s is required", fieldName)
	}

	// Trim whitespace
	name = strings.TrimSpace(name)

	// Check length
	if len(name) < MinNameLength {
		return fmt.Errorf("%s must be at least %d characters", fieldName, MinNameLength)
	}

	if len(name) > MaxNameLength {
		return fmt.Errorf("%s must not exceed %d characters", fieldName, MaxNameLength)
	}

	// Check for valid characters (Unicode letters, spaces, hyphens, apostrophes)
	if !nameRegex.MatchString(name) {
		return fmt.Errorf("%s must contain only letters, spaces, hyphens, and apostrophes", fieldName)
	}

	// Check for numbers
	for _, r := range name {
		if unicode.IsDigit(r) {
			return fmt.Errorf("%s must not contain numbers", fieldName)
		}
	}

	// Check for excessive spaces or hyphens
	if strings.Contains(name, "  ") {
		return fmt.Errorf("%s must not contain consecutive spaces", fieldName)
	}

	if strings.Contains(name, "--") {
		return fmt.Errorf("%s must not contain consecutive hyphens", fieldName)
	}

	return nil
}

// ValidatePersonName validates person's first name
func ValidatePersonName(name string) error {
	return ValidateName(name, "name")
}

// ValidateSurname validates person's surname
func ValidateSurname(surname string) error {
	return ValidateName(surname, "surname")
}

// IsValidName checks if name is valid (returns bool)
func IsValidName(name string) bool {
	return ValidateName(name, "name") == nil
}
