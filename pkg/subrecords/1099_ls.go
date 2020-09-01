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

type Sub1099LS struct {
	// Enter the Date of Sale in format YYYYMMDD (for example
	// January 5, 2019, would be 20190105). Do not enter hyphens
	// or slashes.
	DateSale time.Time `json:"date_sale"`

	// Enter Issuer’s Contact Name.
	IssuersInformation string `json:"issuers_information"`
}

// Type returns type of “1099-LS” record
func (r *Sub1099LS) Type() string {
	return config.Sub1099LsType
}

// Parse parses the “1099-LS” record from fire ascii
func (r *Sub1099LS) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099LSLayout, record)
}

// Ascii returns fire ascii of “1099-LS” record
func (r *Sub1099LS) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099LSLayout)
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
func (r *Sub1099LS) Validate() error {
	return utils.Validate(r, config.Sub1099LSLayout, config.Sub1099LsType)
}
