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

type Sub1098F struct {
	// Enter the effective date of the order in YYYYMMDD format
	// (for example, January 5, 2019, would be 20190105).
	DateOrderAgreement time.Time `json:"date_order_agreement"`

	// Enter the jurisdiction for the fines, penalties, or other
	// amounts being assessed, if applicable.
	Jurisdiction string `json:"jurisdiction"`

	// Enter the case number assigned to the order or agreement, if
	// applicable.
	CaseNumber string `json:"case_number"`

	// Enter a name or description to identify order or agreement.
	MatterSuitAgreement string `json:"matter_suit_agreement"`

	// Enter one or more of the following payment codes.
	// B: Multiple payers/defendants
	// C: Multiple payees
	// D: Property included in settlement
	// E: Settlement payments to nongovernmental entities, i.e., charities
	// F: Settlement paid in full as of time of filing
	// G: No payment received as of time of filing
	// H: Deferred prosecution agreement
	PaymentCode string `json:"payment_code"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or
	// local revenue departments for the filing requirements. If this
	// field is not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “1098-F” record
func (r *Sub1098F) Type() string {
	return config.Sub1098FType
}

// Parse parses the “1098-F” record from fire ascii
func (r *Sub1098F) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1098FLayout, record)
}

// Ascii returns fire ascii of “1098-F” record
func (r *Sub1098F) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1098FLayout)
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
func (r *Sub1098F) Validate() error {
	return utils.Validate(r, config.Sub1098FLayout)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1098F) ValidatePaymentCode() error {
	if !utils.CheckAvailableCodes(r.PaymentCode, config.PaymentCodes1098F) {
		return utils.NewErrValidValue("payment code")
	}
	return nil
}
