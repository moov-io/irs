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

type Sub1099NEC struct {
	// Enter “2” (two) to indicate notification by the IRS twice within
	// three calendar years that the payee provided an incorrect
	// name and/or TIN combination. Otherwise, enter a blank.
	SecondTinNotice string `json:"second_tin_notice"`

	// Enter “1” (one) to indicate sales of $5,000 or more of
	// consumer products to a person on a buy-sell, depositcommission, or any other commission basis for resale
	// anywhere other than in a permanent retail establishment.
	// Otherwise, enter a blank.
	// Note: If reporting a direct sales indicator only, use Type of
	// Return “NE” in Field Positions 26-27, and Amount Code 1
	// in Field Position 28 of the Issuer “A” Record. All payment
	// amount fields in the Payee “B” Record will contain zeros.
	DirectSalesIndicator string `json:"direct_sales_indicator"`

	// State income tax withheld is for the convenience of the filers.
	// This information does not need to be reported to the IRS. The
	// payment amount must be right justified, and unused positions
	// must be zero-filed. If not reporting state income tax withheld,
	// this field may be used as a continuation of the Special Data
	// Entries field.
	StateIncomeTaxWithheld int `json:"state_income_tax_withheld"`

	// Local income tax withheld is for the convenience of the filers.
	// This information does not need to be reported to the IRS. The
	// payment amount must be right justified, and unused positions
	// must be zero-filled. If not reporting local tax withheld, this field
	// may be used as a continuation of the Special Data Entries
	// Field.
	LocalIncomeTaxWithheld int `json:"local_income_tax_withheld"`

	// Enter the valid CF/SF Code if this payee record is to be
	// forwarded toa state agency as part of the CF/SF Program.
	// Enter the valid state code from Part A. Sec. 12, Table 1,
	// Participating States and Codes. Enter Blanks for issuers or
	// states not participating in this program.
	CombinedFSCode int `json:"combined_federal_state_code"`
}

// Type returns type of “1099-NEC” record
func (r *Sub1099NEC) Type() string {
	return config.Sub1099NecType
}

// Type returns FS code of “1099-NEC” record
func (r *Sub1099NEC) FederalState() int {
	return r.CombinedFSCode
}

// Parse parses the “1099-NEC” record from fire ascii
func (r *Sub1099NEC) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099NECLayout, record)
}

// Ascii returns fire ascii of “1099-NEC” record
func (r *Sub1099NEC) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099NECLayout)
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
func (r *Sub1099NEC) Validate() error {
	return utils.Validate(r, config.Sub1099NECLayout, config.Sub1099NecType)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1099NEC) ValidateDirectSalesIndicator() error {
	if len(r.DirectSalesIndicator) > 0 &&
		r.DirectSalesIndicator != config.DirectSalesIndicator {
		return utils.NewErrValidValue("direct sales indicator")
	}
	return nil
}

func (r *Sub1099NEC) ValidateCombinedFSCode() error {
	return utils.ValidateCombinedFSCode(r.CombinedFSCode)
}
