package validator

import (
	"fmt"
	"time"
)

const (
	// MaxAgeYears is the maximum allowed age in years
	MaxAgeYears = 150
)

// ValidateDate validates that a date is in the past and age doesn't exceed limit
func ValidateDate(date time.Time, fieldName string) error {
	if date.IsZero() {
		return fmt.Errorf("%s is required", fieldName)
	}

	now := time.Now()

	// Check if date is in the future
	if date.After(now) {
		return fmt.Errorf("%s cannot be in the future", fieldName)
	}

	// Check if age exceeds maximum
	age := CalculateAge(date)
	if age > MaxAgeYears {
		return fmt.Errorf("%s indicates age exceeds maximum of %d years", fieldName, MaxAgeYears)
	}

	return nil
}

// ValidateBirthdate validates birthdate specifically
func ValidateBirthdate(birthdate time.Time) error {
	return ValidateDate(birthdate, "birthdate")
}

// CalculateAge calculates age in years from birthdate
func CalculateAge(birthdate time.Time) int {
	now := time.Now()

	age := now.Year() - birthdate.Year()

	// Adjust if birthday hasn't occurred this year yet
	if now.Month() < birthdate.Month() ||
		(now.Month() == birthdate.Month() && now.Day() < birthdate.Day()) {
		age--
	}

	return age
}

// IsValidBirthdate checks if birthdate is valid (returns bool)
func IsValidBirthdate(birthdate time.Time) bool {
	return ValidateBirthdate(birthdate) == nil
}
