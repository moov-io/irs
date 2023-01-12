// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"bytes"
	"reflect"
	"unicode/utf8"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/utils"
)

type TRecord struct {
	// Required. Enter “T.”
	RecordType string `json:"record_type" validate:"required"`

	// Required. Enter “2019.”Foreign
	// If reporting prior year data, report the year which applies (2018, 2017, etc.) and set the Prior Year Data Indicator in field position 6.
	PaymentYear int `json:"payment_year" validate:"required"`

	// Required. Enter “P” only if reporting prior year data. Otherwise, enter a blank.
	// Do not enter a “P” if the tax year is 2019.
	// The FIRE System accepts 2010 through 2018 for prior years. You cannot mix tax years within a file.
	PriorYearDataIndicator string `json:"prior_year_data_indicator"`

	// Required. Enter the transmitter’s nine-digit taxpayer identification number (TIN).
	TIN string `json:"transmitter_tin" validate:"required"`

	// Required. Enter the five-character alphanumeric Transmitter Control Code (TCC) assigned by the IRS.
	TCC string `json:"transmitter_control_code" validate:"required"`

	// Required for test files only. Enter a “T” if this is a test file. Otherwise, enter a blank.
	TestFileIndicator string `json:"test_file_indicator"`

	// Enter “1” (one) if the transmitter is a foreign entity. If the transmitter is not a foreign entity, enter a blank.
	ForeignEntityIndicator string `json:"foreign_entity_indicator"`

	// Required. Enter the transmitter name.
	// Left justify the information and fill unused positions with blanks
	TransmitterName string `json:"transmitter_name" validate:"required"`

	// Enter any additional information that may be part of the name.
	// Left justify the information and fill unused positions with blanks.
	TransmitterNameContinuation string `json:"transmitter_name_contd"`

	// Required. Enter company name associated with the address in field positions 190-229.
	CompanyName string `json:"company_name" validate:"required"`

	// Enter any additional information that may be part of the company name.
	CompanyNameContinuation string `json:"company_name_contd"`

	// Required. Enter the mailing address associated with the Company Name in field positions 110-149 where correspondence should be sent.
	// For U.S. address, the payer city, state, and ZIP Code must be reported as a 40-, 2-, and 9-position field, respectively.
	// Filers must adhere to the correct format for the payer city, state, and ZIP Code.
	// For foreign address, filers may use the payer city, state, and ZIP Code as a continuous 51-position field.
	// Enter information in the following order: city, province or state, postal code, and the name of the country.
	// When reporting a foreign address, the Foreign Entity Indicator in position 29 must contain a “1” (one).
	CompanyMailingAddress string `json:"company_mailing_address" validate:"required"`

	// Required. Enter the city, town, or post office where correspondence should be sent.
	CompanyCity string `json:"company_city" validate:"required"`

	// Required. Enter U.S. Postal Service state abbreviation.
	CompanyState string `json:"company_state" validate:"required"`

	// Required. Enter the nine-digit ZIP Code assigned by the U.S.
	// Postal Service. If only the first five digits are known, left justify the information and fill unused positions with blanks.
	CompanyZipCode string `json:"company_zip_code" validate:"required"`

	// Enter the total number of Payee “B” Records reported in the file.
	// Right justify the information and fill unused positions with zeros.
	TotalNumberPayees int `json:"total_number_of_payees"`

	// Required. Enter the name of the person to contact when problems with the file or transmission are encountered.
	ContactName string `json:"contact_name" validate:"required"`

	// Required. Enter the telephone number of the person to contact regarding electronic files. Omit hyphens.
	// If no extension is available, left justify the information and fill unused positions with blanks.
	// Example: The IRS telephone number of 866-455-7438 with an extension of 52345 would be 866455743852345.
	ContactTelephoneNumber string `json:"contact_telephone_number_and_ext" validate:"required"`

	// Required if available. Enter the email address of the person to contact regarding electronic files.
	// If no email address is available, enter blanks. Left justify.
	ContactEmailAddress string `json:"contact_email_address"`

	// Required. Enter the number of the record as it appears within the
	// file. The record sequence number for the “T” Record will always be
	// one (1) since it is the first record on the file and the file can have
	// only one “T” Record. Each record thereafter must be increased by
	// one in ascending numerical sequence, that is, 2, 3, 4, etc. Right
	// justify numbers with leading zeros in the field. For example, the “T”
	// Record sequence number would appear as “00000001” in the field,
	// the first “A” Record would be “00000002,” the first “B” Record,
	// “00000003,” the second “B” Record, “00000004” and so on through
	// the final record of the file, the “F” Record.
	RecordSequenceNumber int `json:"record_sequence_number" validate:"required"`

	// Required. If the software used to produce this file was provided by
	// a vendor or produced in-house, enter the appropriate code from the
	// table below.
	// V: Software was purchased from a vendor or other source.
	// I: Software was produced by in-house programmers.
	VendorIndicator string `json:"vendor_indicator" validate:"required"`

	// Required. Enter the name of the company from whom the software
	// was purchased. If the software is produced in-house, enter blanks
	VendorName string `json:"vendor_name" validate:"required"`

	// Required. Enter the mailing address. If the software is produced
	// in-house, enter blanks.
	// For U.S. address, the payer city, state, and ZIP Code must be
	// reported as a 40-, 2-, and 9-position field, respectively. Filers must
	// adhere to the correct format for the payer city, state, and ZIP Code.
	// For foreign address, filers may use the payer city, state, and ZIP
	// Code as a continuous 51-position field. Enter information in the
	// following order: city, province or state, postal code, and the name of
	// the country. When reporting a foreign address, the Foreign Entity
	// Indicator in position 29 must contain a “1” (one).
	VendorMailingAddress string `json:"vendor_mailing_address" validate:"required"`

	// Required. Enter the city, town, or post office. If the software is
	// produced in-house, enter blanks.
	VendorCity string `json:"vendor_city" validate:"required"`

	// Required. Enter U.S. Postal Service state abbreviation.
	VendorState string `json:"vendor_state" validate:"required"`

	// Required. Enter the valid nine-digit ZIP Code assigned by the U.S.
	// Postal Service. If only the first five digits are known, fill unused
	// positions with blanks. Left justify. If the software is produced inhouse, enter blanks.
	VendorZipCode string `json:"vendor_zip_code" validate:"required"`

	// Required. Enter the name of the person to contact concerning
	// software questions. If the software is produced in-house, enter
	// blanks.
	VendorContactName string `json:"vendor_contact_name" validate:"required"`

	// Required. Enter the telephone number of the person to contact
	// concerning software questions. Omit hyphens. If no extension is
	// available, left justify the information and fill unused positions with
	// blanks. If the software is produced in-house, enter blanks.
	VendorContactTelephoneNumber string `json:"vendor_contact_telephone_and_ext" validate:"required"`

	// Enter “1” (one) if the vendor is a foreign entity. Otherwise, enter a blank.
	VendorForeignEntityIndicator string `json:"vendor_foreign_entity_indicator"`
}

// Type returns type of “T” record
func (r *TRecord) Type() string {
	return r.RecordType
}

// Parse parses the “T” record from fire ascii
func (r *TRecord) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.RecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.TRecordLayout, record)
}

// Ascii returns fire ascii of “T” record
func (r *TRecord) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.TRecordLayout)
	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return nil
	}

	buf.Grow(config.RecordLength)
	for _, spec := range records {
		value := utils.ToString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.Bytes()
}

// Validate performs some checks on the record and returns an error if not Validated
func (r *TRecord) Validate() error {
	return utils.Validate(r, config.TRecordLayout, config.TRecordType)
}

// SequenceNumber returns sequence number of the record
func (r *TRecord) SequenceNumber() int {
	return r.RecordSequenceNumber
}

// SequenceNumber set sequence number of the record
func (r *TRecord) SetSequenceNumber(number int) {
	r.RecordSequenceNumber = number
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *TRecord) ValidateRecordSequenceNumber() error {
	if r.RecordSequenceNumber < 1 {
		return utils.NewErrValidValue("sequence number")
	}
	return nil
}

func (r *TRecord) ValidatePriorYearDataIndicator() error {
	// “P” or a blank.
	if r.PriorYearDataIndicator == config.PriorYearDataIndicator || len(r.PriorYearDataIndicator) == 0 {
		return nil
	}
	return utils.NewErrValidValue("prior year data indicator")
}

func (r *TRecord) ValidateTestFileIndicator() error {
	if r.TestFileIndicator == config.TestFileIndicator || len(r.TestFileIndicator) == 0 {
		return nil
	}
	return utils.NewErrValidValue("test file indicator")
}

func (r *TRecord) ValidateForeignEntityIndicator() error {
	if r.ForeignEntityIndicator == config.ForeignEntityIndicator || len(r.ForeignEntityIndicator) == 0 {
		return nil
	}
	return utils.NewErrValidValue("foreign entity indicator")
}

func (r *TRecord) ValidateCompanyState() error {
	if _, ok := config.StateAbbreviationCodes[r.CompanyState]; ok {
		return nil
	}
	return utils.NewErrValidValue("company state")
}

func (r *TRecord) ValidateVendorIndicator() error {
	if r.VendorIndicator == config.VendorIndicatorProduced || r.VendorIndicator == config.VendorIndicatorPurchased {
		return nil
	}
	return utils.NewErrValidValue("vendor indicator")
}

func (r *TRecord) ValidateVendorState() error {
	if _, ok := config.StateAbbreviationCodes[r.VendorState]; ok {
		return nil
	}
	return utils.NewErrValidValue("vendor state")
}

func (r *TRecord) ValidateVendorForeignEntityIndicator() error {
	if r.VendorForeignEntityIndicator == config.ForeignEntityIndicator || len(r.VendorForeignEntityIndicator) == 0 {
		return nil
	}
	return utils.NewErrValidValue("vendor foreign entity indicator")
}

func (r *TRecord) ValidateCompanyZipCode() error {
	if len(r.CompanyZipCode) >= 0 {
		return utils.IsNumeric(r.CompanyZipCode)
	}
	return utils.NewErrValidValue("company zip code")
}

func (r *TRecord) ValidateVendorZipCode() error {
	if len(r.VendorZipCode) >= 0 {
		return utils.IsNumeric(r.VendorZipCode)
	}
	return utils.NewErrValidValue("vendor zip code")
}
