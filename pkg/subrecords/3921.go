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

type Sub3921 struct {
	// Required. Enter the date the option was granted in
	// YYYYMMDD format (for example, January 5, 2019, would be
	// 20190105).
	DateOptionGranted time.Time `json:"date_option_granted"`

	// Required. Enter the date the option was exercised in
	// YYYYMMDD format (for example, January 5, 2019, would be
	// 20190105).
	DateOptionExercised time.Time `json:"date_option_exercised"`

	// Required. Enter the number of shares transferred. Report
	// whole numbers only, using standard rounding rules as
	// necessary. Right justify the information and fill unused
	// positions with zeros.
	NumberSharesTransferred int `json:"number_shares_transferred"`

	// Enter other than transferor information, left justify the
	// information and fill unused positions with blanks.
	OtherThanTransferorInformation string `json:"other_than_transferor_information"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for filing requirements.
	// If this field is not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “3921” record
func (r *Sub3921) Type() string {
	return config.Sub3921Type
}

// Type returns FS code of “3921” record
func (r *Sub3921) FederalState() int {
	return 0
}

// Parse parses the “3921” record from fire ascii
func (r *Sub3921) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub3921Layout, record)
}

// Ascii returns fire ascii of “3921” record
func (r *Sub3921) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub3921Layout)
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
func (r *Sub3921) Validate() error {
	return utils.Validate(r, config.Sub3921Layout, config.Sub3921Type)
}
