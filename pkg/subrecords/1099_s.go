// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package subrecords

import (
	"bytes"
	"reflect"
	"unicode/utf8"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/utils"
)

type Sub1099S struct {
	// Enter “1” (one) only if the taxable amount of the payment
	// entered for Payment Amount Field 1 (Gross distribution) of
	// the “B” Record cannot be computed. Otherwise, enter a
	// blank. (If the Taxable Amount Not Determined Indicator is
	// used, enter “0s” [zeros] in Payment Amount Field 2 of the
	// Payee “B” Record.) Please make every effort to compute the
	// taxable amount.
	PropertyServicesIndicator string `json:"property_services_indicator"`

	// Required. Enter the closing date in YYYYMMDD format (for
	// example, January 5, 2019, would be 20190105). Do not
	// enter hyphens or slashes.
	DateClosing string `json:"date_closing"`

	// Required. Enter the address of the property transferred
	// (including city, state, and ZIP Code). If the address does not
	// sufficiently identify the property, also enter a legal
	// description, such as section, lot, and block. For timber
	// royalties, enter “TIMBER.”
	// If fewer than 39 positions are required, left justify the
	// information and fill unused positions with blanks.
	AddressLegalDescription string `json:"address_legal_description"`

	// Required. Enter “1” if the transferor is a foreign person
	//(nonresident alien, foreign partnership, foreign estate, or
	//foreign trust). Otherwise, enter a blank.
	ForeignTransferor string `json:"foreign_transferor"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for filing requirements. You may enter
	// your routing and transit number (RTN) here. If this field is not
	// used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`

	// State income tax withheld is for the convenience of the filers.
	// This information does not need to be reported to the IRS. If
	// not reporting state tax withheld, this field may be used as a
	// continuation of the Special Data Entries Field. The payment
	// amount must be right justified and unused positions
	// zero-filled.
	StateIncomeTaxWithheld int `json:"state_income_tax_withheld"`

	// Local income tax withheld is for the convenience of the filers.
	// This information does not need to be reported to the IRS. If
	// not reporting local tax withheld, this field may be used as a
	// continuation of the Special Data Entries Field. The payment
	// amount must be right justified and unused positions
	// zero-filled.
	LocalIncomeTaxWithheld int `json:"local_income_tax_withheld"`
}

// Type returns type of “1099-S” record
func (r *Sub1099S) Type() string {
	return config.Sub1099SType
}

// Parse parses the “1099-S” record from fire ascii
func (r *Sub1099S) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099SLayout, record)
}

// Ascii returns fire ascii of “1099-S” record
func (r *Sub1099S) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099SLayout)
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
func (r *Sub1099S) Validate() error {
	return utils.Validate(r, config.Sub1099SLayout)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1099S) ValidatePropertyServicesIndicator() error {
	if len(r.PropertyServicesIndicator) > 0 &&
		r.PropertyServicesIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("property services indicator")
	}
	return nil
}

func (r *Sub1099S) ValidateForeignTransferor() error {
	if len(r.ForeignTransferor) > 0 &&
		r.ForeignTransferor != config.GeneralOneIndicator {
		return utils.NewErrValidValue("foreign transferor")
	}
	return nil
}
