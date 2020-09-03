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

type Sub1099A struct {
	// Enter the appropriate indicator from the table below:
	// 1: Borrower was personally liable for repayment of the debt.
	// Blank: Borrower was not personally liable for repayment of the debt.
	PersonalLiabilityIndicator string `json:"personal_liability_indicator"`

	// Enter the acquisition date of the secured property or the date
	// the lender first knew or had reason to know the property was
	// abandoned, in YYYYMMDD format (for example, January 5,
	// 2019, would be 20190105). Do not enter hyphens or slashes.
	DateAcquisitionKnowledgeAbandonment time.Time `json:"date_acquisition_knowledge_abandonment"`

	// Enter a brief description of the property. For real property,
	// enter the address, or if the address does not sufficiently
	// identify the property, enter the section, lot and block. For
	// personal property, enter the type, make and model (for
	// example, Car-1999 Buick Regal or Office Equipment). Enter
	// “CCC” for crops forfeited on Commodity Credit Corporation
	// loans.
	// If fewer than 39 positions are required, left justify the
	// information and fill unused positions with blanks.
	DescriptionProperty string `json:"description_property"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for the filing requirements. If this field is
	// not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “1099-A” record
func (r *Sub1099A) Type() string {
	return config.Sub1099AType
}

// Type returns FS code of “1099-A” record
func (r *Sub1099A) FederalState() int {
	return 0
}

// Parse parses the “1099-A” record from fire ascii
func (r *Sub1099A) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099ALayout, record)
}

// Ascii returns fire ascii of “1099-A” record
func (r *Sub1099A) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099ALayout)
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
func (r *Sub1099A) Validate() error {
	return utils.Validate(r, config.Sub1099ALayout, config.Sub1099AType)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1099A) ValidatePersonalLiabilityIndicator() error {
	if len(r.PersonalLiabilityIndicator) > 0 &&
		r.PersonalLiabilityIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("personal liability indicator")
	}
	return nil
}
