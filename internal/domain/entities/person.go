package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/BrunnoQ/cadastro-pessoas/pkg/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Sex represents biological sex
type Sex string

const (
	SexMale   Sex = "M"
	SexFemale Sex = "F"
)

// Person represents a person entity
type Person struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Surname   string             `json:"surname" bson:"surname"`
	Birthdate time.Time          `json:"birthdate" bson:"birthdate"`
	Sex       Sex                `json:"sex" bson:"sex"`
	Addresses []Address          `json:"addresses,omitempty" bson:"addresses,omitempty"`
	Contacts  []Contact          `json:"contacts,omitempty" bson:"contacts,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	Version   int                `json:"version" bson:"version"`
}

// NewPerson creates a new Person with validation
func NewPerson(name, surname string, birthdate time.Time, sex Sex) (*Person, error) {
	person := &Person{
		ID:        primitive.NewObjectID(),
		Name:      strings.TrimSpace(name),
		Surname:   strings.TrimSpace(surname),
		Birthdate: birthdate,
		Sex:       sex,
		Addresses: make([]Address, 0),
		Contacts:  make([]Contact, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}

	if err := person.Validate(); err != nil {
		return nil, err
	}

	return person, nil
}

// Validate validates the person fields
func (p *Person) Validate() error {
	// Validate name
	if err := validator.ValidatePersonName(p.Name); err != nil {
		return err
	}

	// Validate surname
	if err := validator.ValidateSurname(p.Surname); err != nil {
		return err
	}

	// Validate birthdate
	if err := validator.ValidateBirthdate(p.Birthdate); err != nil {
		return err
	}

	// Validate sex
	if p.Sex == "" {
		return fmt.Errorf("sex is required")
	}

	// Case-insensitive sex validation
	sexUpper := Sex(strings.ToUpper(string(p.Sex)))
	if sexUpper != SexMale && sexUpper != SexFemale {
		return fmt.Errorf("sex must be 'M' or 'F'")
	}

	// Normalize sex to uppercase
	p.Sex = sexUpper

	return nil
}

// AddAddress adds a new address to the person
func (p *Person) AddAddress(address Address) error {
	if err := address.Validate(); err != nil {
		return fmt.Errorf("invalid address: %w", err)
	}

	p.Addresses = append(p.Addresses, address)
	p.UpdatedAt = time.Now()
	return nil
}

// AddContact adds a new contact to the person
func (p *Person) AddContact(contact Contact) error {
	if err := contact.Validate(); err != nil {
		return fmt.Errorf("invalid contact: %w", err)
	}

	// Check for duplicate contacts
	for _, existingContact := range p.Contacts {
		if existingContact.Type == contact.Type && existingContact.Value == contact.Value {
			return fmt.Errorf("contact already exists")
		}
	}

	p.Contacts = append(p.Contacts, contact)
	p.UpdatedAt = time.Now()
	return nil
}

// IncrementVersion increments the version for optimistic locking
func (p *Person) IncrementVersion() {
	p.Version++
	p.UpdatedAt = time.Now()
}
