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

type Sub5498ESA struct {
	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for filing requirements.
	// If this field is not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “5498-ESA” record
func (r *Sub5498ESA) Type() string {
	return config.Sub5498EsaType
}

// Parse parses the “5498-ESA” record from fire ascii
func (r *Sub5498ESA) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub5498ESALayout, record)
}

// Ascii returns fire ascii of “5498-ESA” record
func (r *Sub5498ESA) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub5498ESALayout)
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
func (r *Sub5498ESA) Validate() error {
	return utils.Validate(r, config.Sub5498ESALayout, config.Sub5498EsaType)
}
