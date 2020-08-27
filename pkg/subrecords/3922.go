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

type Sub3922 struct {
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

	// Required. Enter the date the legal title was transferred by
	// the transferor as YYYYMMDD (for example, January 5,
	// 2019, would be 20190105). Otherwise, enter blanks.
	DateLegalTitleTransferred time.Time `json:"date_legal_title_transferred"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for filing requirements.
	// If this field is not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “3922” record
func (r *Sub3922) Type() string {
	return config.Sub3922Type
}

// Parse parses the “3922” record from fire ascii
func (r *Sub3922) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub3922Layout, record)
}

// Ascii returns fire ascii of “3922” record
func (r *Sub3922) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub3922Layout)
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
func (r *Sub3922) Validate() error {
	return utils.Validate(r, config.Sub3922Layout)
}
