package utils

import (
	"errors"
	"fmt"
)

var (
	// ErrNonAlphanumeric is given when a field has non-alphanumeric characters
	ErrNonAlphanumeric = errors.New("has non alphanumeric characters")
	// ErrFieldRequired is given when a field is required
	ErrFieldRequired = errors.New("is an invalid required field")
	// ErrUpperAlpha is given when a field is not numeric characters
	ErrNumeric = errors.New("is not numeric characters")
	// ErrUpperAlpha is given when a field is an invalid phone number
	ErrPhoneNumber = errors.New("is an invalid phone number")
	// ErrValidYear is given when there's an invalid date
	ErrValidDate = errors.New("is an invalid Date")
	// ErrValidYear is given when a segment has an invalid length
	ErrRecordLength = errors.New("has an invalid length")
	// ErrValidField is given when there's an invalid field
	ErrValidField = errors.New("is an invalid field")
	// ErrShortRecord is given when the record is too short
	ErrShortRecord = errors.New("is too short / missing data")
	// ErrEmail is given when a field is not email
	ErrEmail = errors.New("is not email address")
)

// NewErrValidValue returns a error that has invalid value
func NewErrValidValue(field string) error {
	return fmt.Errorf("is an invalid value of %s", field)
}
