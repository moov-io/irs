// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package subrecords

import (
	"bytes"
	"reflect"
	"time"
	"unicode/utf8"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/utils"
)

type Sub1098 struct {
	// Enter the date of the Mortgage Origination in YYYYMMDD
	// format.
	MortgageOriginationDate time.Time `json:"mortgage_origination_date"`

	// Enter “1” (one) if Property Securing Mortgage is the same as
	// payer/borrowers’ address. Otherwise enter a blank.
	PropertySecuringMortgageIndicator string `json:"property_securing_mortgage_indicator"`

	// Enter the address or description of the property securing the
	// mortgage if different than the payer/borrowers address.
	// Left justify and fill with blanks.
	PropertyADSecuringMortgage string `json:"property_address_description_securing_mortgage"`

	// Enter any other item you wish to report to the payer.
	// Examples include:
	// • Continuation of Property Address Securing Mortgage
	// • Continuation of Legal Description of Property
	// • Real estate taxes
	// • Insurance paid from escrow
	// • If you are a collection agent, the name of the person for
	// whom you collected the interest
	// This is a free format field. If this field is not used, enter
	// blanks.
	// You do not have to report to the IRS any information
	// provided in this box.
	// Left justify and fill with blanks.
	Other string `json:"other"`

	// If more than one property securing the mortgage, enter the
	// total number of properties secured by this mortgage. If less
	// than two (2), enter blanks. Valid values are 0000 - 9999.
	NumberMortgagedProperties int `json:"number_mortgaged_properties"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or
	// local revenue departments for the filing requirements. If
	// this field is not use, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`

	// Enter the date in format YYYYMMDD if the recipient/lender
	// acquired the mortgage in 2019, show the date of
	// acquisition. (for example, January 5, 2019, would be
	// 20190105)
	MortgageAcquisitionDate time.Time `json:"mortgage_acquisition_date"`
}

// Type returns type of “1098” record
func (r *Sub1098) Type() string {
	return config.Sub1098Type
}

// Parse parses the “1098” record from fire ascii
func (r *Sub1098) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1098Layout, record)
}

// Ascii returns fire ascii of “1098” record
func (r *Sub1098) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1098Layout)
	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return nil
	}

	buf.Grow(config.SubRecordLength)
	for _, spec := range records {
		value := utils.ToString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.Bytes()
}

// Validate performs some checks on the record and returns an error if not Validated
func (r *Sub1098) Validate() error {
	return utils.Validate(r, config.Sub1098Layout, config.Sub1098Type)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1098) ValidatePropertySecuringMortgageIndicator() error {
	if len(r.PropertySecuringMortgageIndicator) > 0 &&
		r.PropertySecuringMortgageIndicator != config.PropertySecuringMortgageIndicator {
		return utils.NewErrValidValue("property securing mortgage indicator")
	}
	return nil
}
