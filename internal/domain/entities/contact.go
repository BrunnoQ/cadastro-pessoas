package entities

import (
	"fmt"
	"strings"

	"github.com/BrunnoQ/cadastro-pessoas/pkg/validator"
)

// ContactType represents the type of contact
type ContactType string

const (
	ContactTypeEmail ContactType = "email"
	ContactTypePhone ContactType = "phone"
)

// Contact represents a person's contact information (Value Object)
type Contact struct {
	Type  ContactType `json:"type" bson:"type"`
	Value string      `json:"value" bson:"value"`
}

// NewContact creates a new Contact with validation
func NewContact(contactType ContactType, value string) (*Contact, error) {
	contact := &Contact{
		Type:  contactType,
		Value: strings.TrimSpace(value),
	}

	if err := contact.Validate(); err != nil {
		return nil, err
	}

	return contact, nil
}

// Validate validates the contact fields
func (c *Contact) Validate() error {
	if c.Type == "" {
		return fmt.Errorf("contact type is required")
	}

	if c.Type != ContactTypeEmail && c.Type != ContactTypePhone {
		return fmt.Errorf("contact type must be 'email' or 'phone'")
	}

	if c.Value == "" {
		return fmt.Errorf("contact value is required")
	}

	// Validate based on contact type
	switch c.Type {
	case ContactTypeEmail:
		if err := validator.ValidateEmail(c.Value); err != nil {
			return fmt.Errorf("invalid email: %w", err)
		}
	case ContactTypePhone:
		if err := validator.ValidatePhone(c.Value); err != nil {
			return fmt.Errorf("invalid phone: %w", err)
		}
	}

	return nil
}
