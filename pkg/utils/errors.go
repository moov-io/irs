package utils

import (
	"errors"
	"fmt"
)

var (
	// ErrNonAlphanumeric is given when a field has non-alphanumeric characters
	ErrNonAlphanumeric = errors.New("has non alphanumeric characters")
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
	// ErrPayeeExtBlock is given when payee record has not extension block for  for each type of return
	ErrPayeeExtBlock = errors.New("should exist extension block")
	// ErrInvalidAscii is given when is invalid ascii
	ErrInvalidAscii = errors.New("is invalid ascii")
	// ErrInvalidFile is given when is invalid file
	ErrInvalidFile = errors.New("is invalid file")
)

// NewErrValidValue returns a error that has invalid value
func NewErrValidValue(field string) error {
	return fmt.Errorf("is an invalid value of %s", field)
}

// NewErrFieldRequired returns a error that has empty required field
func NewErrFieldRequired(field string) error {
	return fmt.Errorf("is required field (%s)", field)
}
