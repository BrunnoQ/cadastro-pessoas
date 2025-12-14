package dto

import (
	"github.com/BrunnoQ/cadastro-pessoas/internal/domain/entities"
)

// CreatePersonRequest represents a request to create a person
type CreatePersonRequest struct {
	Name      string           `json:"name"`
	Surname   string           `json:"surname"`
	Birthdate DateOnly         `json:"birthdate"`
	Sex       entities.Sex     `json:"sex"`
	Addresses []AddressRequest `json:"addresses,omitempty"`
	Contacts  []ContactRequest `json:"contacts,omitempty"`
}

// AddressRequest represents an address in a request
type AddressRequest struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement,omitempty"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
}

// ContactRequest represents a contact in a request
type ContactRequest struct {
	Type  entities.ContactType `json:"type"`
	Value string               `json:"value"`
}

// UpdatePersonRequest represents a request to update a person
type UpdatePersonRequest struct {
	Name      *string           `json:"name,omitempty"`
	Surname   *string           `json:"surname,omitempty"`
	Birthdate *DateOnly         `json:"birthdate,omitempty"`
	Sex       *entities.Sex     `json:"sex,omitempty"`
	Addresses *[]AddressRequest `json:"addresses,omitempty"`
	Contacts  *[]ContactRequest `json:"contacts,omitempty"`
}

// ToAddress converts AddressRequest to Address entity
func (r *AddressRequest) ToAddress() (*entities.Address, error) {
	return entities.NewAddress(
		r.Street,
		r.Number,
		r.Complement,
		r.Neighborhood,
		r.City,
		r.State,
		r.Country,
		r.ZipCode,
	)
}

// ToContact converts ContactRequest to Contact entity
func (r *ContactRequest) ToContact() (*entities.Contact, error) {
	return entities.NewContact(r.Type, r.Value)
}
