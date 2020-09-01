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

type Sub1099C struct {
	// Required. Enter the appropriate indicator from the following table:
	// A: Bankruptcy
	// B: Other Judicial Debt Relief
	// C: Statute of limitations or expiration of deficiency period
	// D: Foreclosure election
	// E: Debt relief from probate or similar proceeding
	// F: By agreement
	// G: Creditor’s debt collection policy
	// H: Other actual discharge before identifiable event
	IdentifiableEventCode string `json:"identifiable_event_code"`

	// Enter the date the debt was canceled in YYYYMMDD format
	// (for example, January 5, 2019, would be 20190105). Do not
	// enter hyphens or slashes.
	DateIdentifiableEvent time.Time `json:"date_identifiable_event"`

	// Enter a description of the origin of the debt, such as student
	// loan, mortgage, or credit card expenditure. If a combined
	// Form 1099-C and 1099-A is being filed, also enter a
	// description of the property.
	DebtDescription string `json:"debt_description"`

	// Enter “1” (one) if the borrower is personally liable for
	// repayment, or enter a blank if not personally liable for
	// repayment.
	PersonalLiabilityIndicator string `json:"personal_liability_indicator"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for the filing requirements. If this field is
	// not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “1099-C” record
func (r *Sub1099C) Type() string {
	return config.Sub1099CType
}

// Parse parses the “1099-C” record from fire ascii
func (r *Sub1099C) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099CLayout, record)
}

// Ascii returns fire ascii of “1099-C” record
func (r *Sub1099C) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099CLayout)
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
func (r *Sub1099C) Validate() error {
	return utils.Validate(r, config.Sub1099CLayout, config.Sub1099CType)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1099C) ValidateIdentifiableEventCode() error {
	if len(r.IdentifiableEventCode) > 0 {
		switch r.IdentifiableEventCode {
		case "A", "B", "C", "D", "E", "F", "G", "H":
			break
		default:
			return utils.NewErrValidValue("identifiable event code")
		}
	}
	return nil
}

func (r *Sub1099C) ValidatePersonalLiabilityIndicator() error {
	if len(r.PersonalLiabilityIndicator) > 0 &&
		r.PersonalLiabilityIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("personal liability indicator")
	}
	return nil
}
