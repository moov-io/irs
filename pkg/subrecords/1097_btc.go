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

type Sub1097BTC struct {
	// Required. Enter the appropriate indicator from the table
	// below:
	// 1: Issuer of bond or its agent filing
	//    initial 2019 Form 1097-BTC for
	//    credit being reported
	// 2: An entity that received a 2018 Form
	//    1097-BTC for credit being reported
	IssuerIndicator string `json:"issuer_indicator" validate:"required"`

	// Required. Enter the appropriate alpha indicator from the
	// table below:
	// A: Account number
	// B: CUSIP number
	// O: Unique identification number, not an
	//    account/CUSIP number, such as a
	//    self-provided identification number
	Code string `json:"code" validate:"required"`

	// Enter the unique identifier assigned to the bond. This can be
	// an alphanumeric identifier such as the CUSIP number.
	// Right justify the information and fill unused positions with
	// blanks.
	UniqueIdentifier string `json:"unique_identifier"`

	// Required. Enter the appropriate indicator from the table.
	// 101: Clean Renewable Energy Bond
	// 199: Other
	BondType string `json:"bond_type" validate:"required"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for filing requirements. You may enter
	// comments here. If this field is not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “1097-BTC” record
func (r *Sub1097BTC) Type() string {
	return config.Sub1097BtcType
}

// Parse parses the “1097-BTC” record from fire ascii
func (r *Sub1097BTC) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1097BTCLayout, record)
}

// Ascii returns fire ascii of “1097-BTC” record
func (r *Sub1097BTC) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1097BTCLayout)
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
func (r *Sub1097BTC) Validate() error {
	return utils.Validate(r, config.Sub1097BTCLayout)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1097BTC) ValidateIssuerIndicator() error {
	if _, ok := config.BtcIssuerIndicator[r.IssuerIndicator]; !ok {
		return utils.NewErrValidValue("issuer indicator")
	}
	return nil
}

func (r *Sub1097BTC) ValidateCode() error {
	if _, ok := config.BtcCode[r.Code]; !ok {
		return utils.NewErrValidValue("code")
	}
	return nil
}

func (r *Sub1097BTC) ValidateBondType() error {
	if _, ok := config.BtcBondType[r.BondType]; !ok {
		return utils.NewErrValidValue("bond type")
	}
	return nil
}
