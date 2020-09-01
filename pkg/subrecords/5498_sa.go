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

type Sub5498SA struct {
	// Enter “1” (one) for a Medicare Advantage MSA. Otherwise,
	// enter a blank.
	MedicareAdvantageMSAIndicator string `json:"medicare_advantage_msa_indicator"`

	// Enter “1” (one) for an HSA. Otherwise, enter a blank.
	HSAIndicator string `json:"hsa_indicator"`

	// Enter “1” (one) for an Archer MSA. Otherwise, enter a blank.
	ArcherMSAIndicator string `json:"archer_mas_indicator"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for filing requirements.
	// If this field is not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “5498-SA” record
func (r *Sub5498SA) Type() string {
	return config.Sub5498SaType
}

// Parse parses the “5498-SA” record from fire ascii
func (r *Sub5498SA) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub5498SALayout, record)
}

// Ascii returns fire ascii of “5498-SA” record
func (r *Sub5498SA) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub5498SALayout)
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
func (r *Sub5498SA) Validate() error {
	return utils.Validate(r, config.Sub5498SALayout, config.Sub5498SaType)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub5498SA) ValidateMedicareAdvantageMSAIndicator() error {
	if len(r.MedicareAdvantageMSAIndicator) > 0 &&
		r.MedicareAdvantageMSAIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("medicare advantage msa indicator")
	}
	return nil
}

func (r *Sub5498SA) ValidateHSAIndicator() error {
	if len(r.HSAIndicator) > 0 &&
		r.HSAIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("has indicator")
	}
	return nil
}

func (r *Sub5498SA) ValidateArcherMSAIndicator() error {
	if len(r.ArcherMSAIndicator) > 0 &&
		r.ArcherMSAIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("archer mas indicator")
	}
	return nil
}
