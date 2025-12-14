package entities

import (
	"fmt"
	"strings"

	"github.com/BrunnoQ/cadastro-pessoas/pkg/errors"
	"github.com/BrunnoQ/cadastro-pessoas/pkg/validator"
)

// Address represents a physical address
type Address struct {
	Street       string `json:"street" bson:"street"`
	Number       string `json:"number" bson:"number"`
	Complement   string `json:"complement,omitempty" bson:"complement,omitempty"`
	Neighborhood string `json:"neighborhood" bson:"neighborhood"`
	City         string `json:"city" bson:"city"`
	State        string `json:"state" bson:"state"`
	Country      string `json:"country" bson:"country"`
	ZipCode      string `json:"zip_code" bson:"zip_code"`
}

// NewAddress creates a new Address with trimmed fields
func NewAddress(
	street, number, complement, neighborhood, city, state, country, zipCode string,
) (*Address, error) {
	address := &Address{
		Street:       strings.TrimSpace(street),
		Number:       strings.TrimSpace(number),
		Complement:   strings.TrimSpace(complement),
		Neighborhood: strings.TrimSpace(neighborhood),
		City:         strings.TrimSpace(city),
		State:        strings.TrimSpace(state),
		Country:      strings.TrimSpace(country),
		ZipCode:      strings.TrimSpace(zipCode),
	}

	if err := address.Validate(); err != nil {
		return nil, err
	}

	return address, nil
}

// Validate validates the address fields
func (a *Address) Validate() error {
	// Street is required
	if a.Street == "" {
		return errors.NewValidationError("street is required", map[string]interface{}{
			"field": "street",
		})
	}
	if len(a.Street) > 200 {
		return errors.NewValidationError("street must not exceed 200 characters", map[string]interface{}{
			"field":      "street",
			"max_length": 200,
		})
	}

	// Number is required
	if a.Number == "" {
		return errors.NewValidationError("number is required", map[string]interface{}{
			"field": "number",
		})
	}
	if len(a.Number) > 20 {
		return errors.NewValidationError("number must not exceed 20 characters", map[string]interface{}{
			"field":      "number",
			"max_length": 20,
		})
	}

	// Complement is optional but has max length
	if len(a.Complement) > 100 {
		return errors.NewValidationError("complement must not exceed 100 characters", map[string]interface{}{
			"field":      "complement",
			"max_length": 100,
		})
	}

	// Neighborhood is required
	if a.Neighborhood == "" {
		return errors.NewValidationError("neighborhood is required", map[string]interface{}{
			"field": "neighborhood",
		})
	}
	if len(a.Neighborhood) > 100 {
		return errors.NewValidationError("neighborhood must not exceed 100 characters", map[string]interface{}{
			"field":      "neighborhood",
			"max_length": 100,
		})
	}

	// City is required
	if a.City == "" {
		return errors.NewValidationError("city is required", map[string]interface{}{
			"field": "city",
		})
	}
	if len(a.City) > 100 {
		return errors.NewValidationError("city must not exceed 100 characters", map[string]interface{}{
			"field":      "city",
			"max_length": 100,
		})
	}
	if err := validator.ValidateName(a.City, "city"); err != nil {
		return fmt.Errorf("invalid city: %w", err)
	}

	// State is required
	if a.State == "" {
		return errors.NewValidationError("state is required", map[string]interface{}{
			"field": "state",
		})
	}
	if len(a.State) > 100 {
		return errors.NewValidationError("state must not exceed 100 characters", map[string]interface{}{
			"field":      "state",
			"max_length": 100,
		})
	}
	if err := validator.ValidateName(a.State, "state"); err != nil {
		return fmt.Errorf("invalid state: %w", err)
	}

	// Country is required
	if a.Country == "" {
		return errors.NewValidationError("country is required", map[string]interface{}{
			"field": "country",
		})
	}
	if len(a.Country) > 100 {
		return errors.NewValidationError("country must not exceed 100 characters", map[string]interface{}{
			"field":      "country",
			"max_length": 100,
		})
	}
	if err := validator.ValidateName(a.Country, "country"); err != nil {
		return fmt.Errorf("invalid country: %w", err)
	}

	// ZipCode is required
	if a.ZipCode == "" {
		return errors.NewValidationError("zip_code is required", map[string]interface{}{
			"field": "zip_code",
		})
	}
	if len(a.ZipCode) > 20 {
		return errors.NewValidationError("zip_code must not exceed 20 characters", map[string]interface{}{
			"field":      "zip_code",
			"max_length": 20,
		})
	}

	return nil
}
