// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"bytes"
	"reflect"
	"unicode/utf8"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/utils"
)

type KRecord struct {
	// Required. Enter “K.”
	RecordType string `json:"record_type" validate:"required"`

	// Required. Enter the total number of “B” Records being
	// coded for this state. Right justify the information and fill
	// unused positions with zeros.
	NumberPayees int `json:"number_of_payees" validate:"required"`

	// Required. Accumulate totals of any payment amount fields
	// in the “B” Records for each state being reported into the
	// appropriate control total fields of the appropriate “K” Record.
	// Each payment amount must contain U.S. dollars and cents.
	// The right-most two positions represent cents in the payment
	// amount fields. Control totals must be right justified and fill
	// unused positions with zeros. All control total fields are
	// eighteen positions in length. Do not enter dollar signs,
	// commas, decimal points, or negative payments, except those
	// items that reflect a loss on Form 1099-B or 1099-OID.
	// Positive and negative amounts are indicated by placing a “+”
	// (plus) or “-” (minus) sign in the left-most position of the
	// payment amount field.
	ControlTotal1 int `json:"control_total_1"`
	ControlTotal2 int `json:"control_total_2"`
	ControlTotal3 int `json:"control_total_3"`
	ControlTotal4 int `json:"control_total_4"`
	ControlTotal5 int `json:"control_total_5"`
	ControlTotal6 int `json:"control_total_6"`
	ControlTotal7 int `json:"control_total_7"`
	ControlTotal8 int `json:"control_total_8"`
	ControlTotal9 int `json:"control_total_9"`
	ControlTotalA int `json:"control_total_A"`
	ControlTotalB int `json:"control_total_B"`
	ControlTotalC int `json:"control_total_C"`
	ControlTotalD int `json:"control_total_D"`
	ControlTotalE int `json:"control_total_E"`
	ControlTotalF int `json:"control_total_F"`
	ControlTotalG int `json:"control_total_G"`

	// Required. Enter the number of the record as it appears
	// within the file. The record sequence number for the “T”
	// Record will always be “1” (one), since it is the first record on
	// the file and the file can have only one “T” Record in a file.
	// Each record, thereafter, must be increased by one in
	// ascending numerical sequence, that is, 2, 3, 4, etc. Right
	// justify numbers with leading zeros in the field. For example,
	// the “T” Record sequence number would appear as
	// “00000001” in the field, the first “A” Record would be
	// “00000002,” the first “B” Record, “00000003,” the second “B”
	// Record, “00000004” and so on through the final record of the
	// file, the “F” Record.
	RecordSequenceNumber int `json:"record_sequence_number" validate:"required"`

	// Aggregate totals of the state income tax withheld field in the
	// Payee “B” Records. Otherwise, enter blanks. (This field is for
	// the convenience of filers.)
	StateIncomeTaxWithheldTotal string `json:"state_income_tax_withheld_total"`

	// Aggregate totals of the local income tax withheld field in the
	// Payee “B” Records. Otherwise, enter blanks. (This field is for
	// the convenience of filers.)
	LocalIncomeTaxWithheldTotal string `json:"local_income_tax_withheld_total"`

	// Required. Enter the CF/SF code assigned to the state which
	// is to receive the information.
	CombinedFederalStateCode string `json:"combined_federal_state_code" validate:"required"`
}

// Type returns type of “K” record
func (r *KRecord) Type() string {
	return config.KRecordType
}

// Parse parses the “K” record from fire ascii
func (r *KRecord) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.RecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.KRecordLayout, record)
}

// Ascii returns fire ascii of “K” record
func (r *KRecord) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.KRecordLayout)
	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return nil
	}

	buf.Grow(config.RecordLength)
	for _, spec := range records {
		value := utils.ToString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.Bytes()
}

// Validate performs some checks on the record and returns an error if not Validated
func (r *KRecord) Validate() error {
	return utils.Validate(r, config.KRecordLayout)
}

// SequenceNumber returns sequence number of the record
func (r *KRecord) SequenceNumber() int {
	return r.RecordSequenceNumber
}

// SequenceNumber set sequence number of the record
func (r *KRecord) SetSequenceNumber(number int) {
	r.RecordSequenceNumber = number
}

// ControlTotal returns total of any payment amount field
func (r *KRecord) ControlTotal(index string) (int, error) {
	value, err := utils.GetField(r, "ControlTotal"+index)
	if err != nil {
		return 0, err
	}
	return int(value.Int()), nil
}

// PaymentAmount returns payment codes
func (r *KRecord) PaymentCodes() string {
	codes := ""
	codeIndexes := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G"}
	for _, index := range codeIndexes {
		amount, err := r.ControlTotal(index)
		if err == nil && amount > 0 {
			codes += index
		}
	}
	return codes
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *KRecord) ValidateRecordSequenceNumber() error {
	if r.RecordSequenceNumber < 1 {
		return utils.NewErrValidValue("sequence number")
	}
	return nil
}

func (r *KRecord) ValidateCombinedFederalStateCode() error {
	if _, ok := config.StateAbbreviationCodes[r.CombinedFederalStateCode]; ok {
		return nil
	}
	return utils.NewErrValidValue("combined federal state code")
}
