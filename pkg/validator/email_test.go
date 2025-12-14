package validator

import (
	"testing"
)

func TestValidateEmail_ValidEmails_ReturnsNil(t *testing.T) {
	validEmails := []string{
		"user@example.com",
		"test.user@example.com",
		"user+tag@example.com",
		"user_name@example.co.uk",
		"first.last@subdomain.example.com",
		"123@example.com",
		"user@example123.com",
		"a@b.co",
	}

	for _, email := range validEmails {
		t.Run(email, func(t *testing.T) {
			err := ValidateEmail(email)
			if err != nil {
				t.Errorf("ValidateEmail(%q) returned error: %v, expected nil", email, err)
			}
		})
	}
}

func TestValidateEmail_EmptyEmail_ReturnsError(t *testing.T) {
	err := ValidateEmail("")
	if err == nil {
		t.Error("ValidateEmail(\"\") expected error, got nil")
	}

	expectedMsg := "email is required"
	if err.Error() != expectedMsg {
		t.Errorf("ValidateEmail(\"\") error = %q, want %q", err.Error(), expectedMsg)
	}
}

func TestValidateEmail_InvalidFormat_ReturnsError(t *testing.T) {
	invalidEmails := []struct {
		email       string
		description string
	}{
		{"invalid", "missing @ symbol"},
		{"@example.com", "missing local part"},
		{"user@", "missing domain"},
		{"user@@example.com", "double @ symbol"},
		{"user@example", "missing TLD"},
	}

	for _, tc := range invalidEmails {
		t.Run(tc.description, func(t *testing.T) {
			err := ValidateEmail(tc.email)
			if err == nil {
				t.Errorf("ValidateEmail(%q) expected error for %s, got nil", tc.email, tc.description)
			}
		})
	}
}

func TestValidateEmail_LocalPartTooLong_ReturnsError(t *testing.T) {
	// Create local part with 65 characters (exceeds max of 64)
	longLocal := ""
	for i := 0; i < 65; i++ {
		longLocal += "a"
	}
	email := longLocal + "@example.com"

	err := ValidateEmail(email)
	if err == nil {
		t.Error("ValidateEmail(email with long local part) expected error, got nil")
	}
}

func TestValidateEmail_TrimWhitespace_Success(t *testing.T) {
	emailWithSpaces := "  user@example.com  "
	err := ValidateEmail(emailWithSpaces)
	if err != nil {
		t.Errorf("ValidateEmail(%q) returned error: %v, expected nil (should trim whitespace)", emailWithSpaces, err)
	}
}

func TestIsValidEmail_ValidEmail_ReturnsTrue(t *testing.T) {
	if !IsValidEmail("user@example.com") {
		t.Error("IsValidEmail(\"user@example.com\") = false, want true")
	}
}

func TestIsValidEmail_InvalidEmail_ReturnsFalse(t *testing.T) {
	if IsValidEmail("invalid") {
		t.Error("IsValidEmail(\"invalid\") = true, want false")
	}
}
