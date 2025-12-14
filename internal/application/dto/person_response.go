package dto

import (
	"time"

	"github.com/BrunnoQ/cadastro-pessoas/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PersonResponse represents a person in a response
type PersonResponse struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Surname   string            `json:"surname"`
	Birthdate time.Time         `json:"birthdate"`
	Sex       entities.Sex      `json:"sex"`
	Addresses []AddressResponse `json:"addresses,omitempty"`
	Contacts  []ContactResponse `json:"contacts,omitempty"`
	Version   int64             `json:"version"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// AddressResponse represents an address in a response
type AddressResponse struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement,omitempty"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
}

// ContactResponse represents a contact in a response
type ContactResponse struct {
	Type  entities.ContactType `json:"type"`
	Value string               `json:"value"`
}

// PersonListResponse represents a list of persons with pagination
type PersonListResponse struct {
	Data    []PersonResponse `json:"data"`
	Total   int64            `json:"total"`
	Limit   int              `json:"limit"`
	Offset  int              `json:"offset"`
	HasMore bool             `json:"has_more"`
}

// FromPerson converts Person entity to PersonResponse
func FromPerson(person *entities.Person) *PersonResponse {
	response := &PersonResponse{
		ID:        person.ID.Hex(),
		Name:      person.Name,
		Surname:   person.Surname,
		Birthdate: person.Birthdate,
		Sex:       person.Sex,
		Version:   int64(person.Version),
		CreatedAt: person.CreatedAt,
		UpdatedAt: person.UpdatedAt,
	}

	// Convert addresses
	if len(person.Addresses) > 0 {
		response.Addresses = make([]AddressResponse, len(person.Addresses))
		for i, addr := range person.Addresses {
			response.Addresses[i] = AddressResponse{
				Street:       addr.Street,
				Number:       addr.Number,
				Complement:   addr.Complement,
				Neighborhood: addr.Neighborhood,
				City:         addr.City,
				State:        addr.State,
				Country:      addr.Country,
				ZipCode:      addr.ZipCode,
			}
		}
	}

	// Convert contacts
	if len(person.Contacts) > 0 {
		response.Contacts = make([]ContactResponse, len(person.Contacts))
		for i, contact := range person.Contacts {
			response.Contacts[i] = ContactResponse{
				Type:  contact.Type,
				Value: contact.Value,
			}
		}
	}

	return response
}

// FromPersons converts multiple Person entities to PersonListResponse
func FromPersons(persons []*entities.Person, total int64, limit, offset int) *PersonListResponse {
	data := make([]PersonResponse, len(persons))
	for i, person := range persons {
		data[i] = *FromPerson(person)
	}

	return &PersonListResponse{
		Data:    data,
		Total:   total,
		Limit:   limit,
		Offset:  offset,
		HasMore: int64(offset+limit) < total,
	}
}

// ParseObjectID parses a string ID to ObjectID
func ParseObjectID(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}
