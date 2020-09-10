// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

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
	// ErrNonExistPayer is given when isn't payer record
	ErrNonExistPayer = errors.New("should exist payer record")
	// ErrNonExistEndPayer is given when isn't payer record
	ErrNonExistEndPayer = errors.New("should exist end of payer record")
	// ErrNonExistPayee is given when isn't payee record
	ErrNonExistPayee = errors.New("should exist at least one payee record")
	// ErrInvalidNumberPayees is given when has incorrect number of payees
	ErrInvalidNumberPayees = errors.New("has incorrect number of payees")
	// ErrIncorrectReturnIndicator is given when has incorrect return indicator
	ErrIncorrectReturnIndicator = errors.New("has incorrect return indicator")
	// ErrInvalidTotalAmounts is given when have invalid totals of any payment amount fields
	ErrInvalidTotalAmounts = errors.New("have invalid totals of any payment amount fields")
	// ErrUnexpectedPaymentAmount is given when has unexpected payment amount in B records
	ErrUnexpectedPaymentAmount = errors.New("has unexpected payment amount")
	// ErrUnexpectedTotalAmount is given when has unexpected totals of any payment amount in C,K record
	ErrUnexpectedTotalAmount = errors.New("has unexpected totals of any payment amount")
	// ErrInvalidTypeOfReturn is given when has invalid type of return
	ErrInvalidTypeOfReturn = errors.New("has invalid type of return")
	// ErrDuplicatedFSCode is given when has duplicated combined fs code in state records
	ErrDuplicatedFSCode = errors.New("has duplicated combined fs code")
	// ErrInvalidNumberPayers is given when has incorrect number of payers
	ErrInvalidNumberPayers = errors.New("has incorrect number of payers")
	// ErrInvalidTCC is given when is invalid transmitter control code
	ErrInvalidTCC = errors.New("is invalid transmitter control code")
	// ErrUnsupportedBlock is given when is not supported extension block of B record
	ErrUnsupportedBlock = errors.New("is not supported extension block of B record")
	// ErrUnknownPdfTemplate is given when is unknown pdf template
	ErrUnknownPdfTemplate = errors.New("is unknown pdf template")
	// ErrFdfGenerate is given when failed to generate fdf file
	ErrFdfGenerate = errors.New("failed to generate fdf file")
	// ErrCFSFProgram is given when has invalid CF/SF program
	ErrCFSFProgram = errors.New("should be payee B records and the state totals K records")
	// ErrCFSFState is given when has invalid Combined Federal/State Code
	ErrCFSFState = errors.New("is invalid combined federal/tate code in K record")
)

// NewErrValidValue returns a error that has invalid value
func NewErrValidValue(field string) error {
	return fmt.Errorf("is an invalid value of %s", field)
}

// NewErrRecordType returns a error that has invalid record type
func NewErrRecordType(field string) error {
	return fmt.Errorf("has invalid record type (%s)", field)
}

// NewErrFieldRequired returns a error that has empty required field
func NewErrFieldRequired(field string) error {
	return fmt.Errorf("is required field (%s)", field)
}

// NewErrRecordSequenceNumber returns a error that has invalid record sequence number
func NewErrRecordSequenceNumber(field string) error {
	return fmt.Errorf("has invalid record sequence number (%s)", field)
}

// NewErrUnexpectedRecord returns a error that has unexpected record
func NewErrUnexpectedRecord(name string, record interface{}) error {
	return fmt.Errorf("unexpected %s record, but got %T", name, record)
}
