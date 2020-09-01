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

type Sub1099H struct {
	// Required. Enter the total number of months the recipient is
	// eligible for health insurance advance payments. Right justify
	// the information and fill unused positions with blanks.
	NumberMonthsEligible string `json:"number_months_eligible"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for the filing requirements. If this field is
	// not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “1099-H” record
func (r *Sub1099H) Type() string {
	return config.Sub1099HType
}

// Parse parses the “1099-H” record from fire ascii
func (r *Sub1099H) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099HLayout, record)
}

// Ascii returns fire ascii of “1099-H” record
func (r *Sub1099H) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099HLayout)
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
func (r *Sub1099H) Validate() error {
	return utils.Validate(r, config.Sub1099HLayout, config.Sub1099HType)
}
