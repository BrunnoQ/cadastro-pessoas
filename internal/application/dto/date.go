package dto

import (
	"fmt"
	"time"
)

const DateFormat = "2006-01-02"

// DateOnly represents a date without time information
type DateOnly struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler
func (d *DateOnly) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	str := string(data)
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}

	// Parse date
	t, err := time.Parse(DateFormat, str)
	if err != nil {
		return fmt.Errorf("invalid date format, expected %s: %w", DateFormat, err)
	}

	d.Time = t
	return nil
}

// MarshalJSON implements json.Marshaler
func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Time.Format(DateFormat))), nil
}
