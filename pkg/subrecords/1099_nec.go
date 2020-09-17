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

	// Enter "1" (one) if there is FATCA filing requirement.
	// Otherwise, enter a blank.
	FATCA string `json:"fatca_requirement_indicator"`
}

// Type returns type of “1099-NEC” record
func (r *Sub1099NEC) Type() string {
	return config.Sub1099NecType
}

// Type returns FS code of “1099-NEC” record
func (r *Sub1099NEC) FederalState() int {
	return 0
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

func (r *Sub1099NEC) ValidateFATCA() error {
	if len(r.FATCA) > 0 &&
		r.FATCA != config.FatcaFilingRequirementIndicator {
		return utils.NewErrValidValue("fatca filing requirement indicator")
	}
	return nil
}
