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

type Sub1099DIV struct {
	// Enter “2” (two) to indicate notification by the IRS twice within
	// three calendar years that the payee provided an incorrect
	// name and/or TIN combination. Otherwise, enter a blank.
	SecondTinNotice string `json:"second_tin_notice"`

	// Enter the name of the foreign country or U.S. possession to
	// which the withheld foreign tax (Amount Code C) applies.
	// Otherwise, enter blanks.
	ForeignCountryPossession string `json:"foreign_country_possession"`

	// Enter "1" (one) if there is FATCA filing requirement.
	// Otherwise, enter a blank.
	FATCA string `json:"fatca_requirement_indicator"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for the filing requirements. If this field is
	// not used, enter blanks.
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

	// Enter the valid CF/SF code if this payee record is to be
	// forwarded to a state agency as part of the CF/SF Program.
	CombinedFSCode int `json:"combined_federal_state_code"`
}

// Type returns type of “1099-DIV” record
func (r *Sub1099DIV) Type() string {
	return config.Sub1099DivType
}

// Parse parses the “1099-DIV” record from fire ascii
func (r *Sub1099DIV) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099DIVLayout, record)
}

// Ascii returns fire ascii of “1099-DIV” record
func (r *Sub1099DIV) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099DIVLayout)
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
func (r *Sub1099DIV) Validate() error {
	return utils.Validate(r, config.Sub1099DIVLayout)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1099DIV) ValidateSecondTinNotice() error {
	if len(r.SecondTinNotice) > 0 &&
		r.SecondTinNotice != config.SecondTINNotice {
		return utils.NewErrValidValue("second tin notice")
	}
	return nil
}

func (r *Sub1099DIV) ValidateFATCA() error {
	if len(r.FATCA) > 0 &&
		r.FATCA != config.FatcaFilingRequirementIndicator {
		return utils.NewErrValidValue("fatca filing requirement indicator")
	}
	return nil
}

func (r *Sub1099DIV) ValidateCombinedFSCode() error {
	if _, ok := config.ParticipateStateCodes[r.CombinedFSCode]; !ok {
		return utils.NewErrValidValue("combined federal state code")
	}
	return nil
}
