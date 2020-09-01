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

type Sub1099SB struct {
	// Enter Issuer’s contact name.
	IssuersInformation string `json:"issuers_information"`
}

// Type returns type of “1099-SB” record
func (r *Sub1099SB) Type() string {
	return config.Sub1099SbType
}

// Parse parses the “1099-SB” record from fire ascii
func (r *Sub1099SB) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099SBLayout, record)
}

// Ascii returns fire ascii of “1099-SB” record
func (r *Sub1099SB) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099SBLayout)
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
func (r *Sub1099SB) Validate() error {
	return utils.Validate(r, config.Sub1099SBLayout, config.Sub1099SbType)
}
