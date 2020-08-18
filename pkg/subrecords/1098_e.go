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

type Sub1098E struct {
	// Enter “1” (one) if the amount reported in Payment Amount
	// Field 1 does not include loan origination fees and/or
	// capitalized interest made before September 1, 2004.
	// Otherwise, enter a blank.
	OriginationInterestIndicator string `json:"origination_interest_indicator"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or
	// local revenue departments for the filing requirements. If
	// this field is not use, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “1098-E” record
func (r *Sub1098E) Type() string {
	return config.Sub1098EType
}

// Parse parses the “1098-E” record from fire ascii
func (r *Sub1098E) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1098ELayout, record)
}

// Ascii returns fire ascii of “1098-E” record
func (r *Sub1098E) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1098ELayout)
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
func (r *Sub1098E) Validate() error {
	return utils.Validate(r, config.Sub1098ELayout)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1098E) ValidateOriginationInterestIndicator() error {
	if len(r.OriginationInterestIndicator) > 0 &&
		r.OriginationInterestIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("origination interest indicator")
	}
	return nil
}
