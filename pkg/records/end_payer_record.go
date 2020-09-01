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

type CRecord struct {
	// Required. Enter “C.”
	RecordType string `json:"record_type" validate:"required"`

	// Required. Enter the total number of “B” Records covered by
	// the preceding “A” Record.
	// Right justify the information and fill unused positions with
	// zeros.
	NumberPayees int `json:"number_of_payees" validate:"required"`

	// Required. Accumulate totals of any payment amount fields
	// in the “B” Records into the appropriate control total fields of
	// the “C” Record. Control totals must be right justified and
	// unused control total fields zero-filled. All control total fields
	// are 18 positions in length. Each payment amount must
	// contain U.S. dollars and cents. The right-most two positions
	// represent cents in the payment amount fields. Do not enter
	// dollar signs, commas, decimal points, or negative payments,
	// except those items that reflect a loss on Form 1099-B, 1099-
	// OID, or 1099-Q. Positive and negative amounts are indicated
	// by placing a “+” (plus) or “-” (minus) sign in the left-most
	// position of the payment amount field.
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
	// Record, “00000004” and so on until the final record of the
	// file, the “F” Record.
	RecordSequenceNumber int `json:"record_sequence_number" validate:"required"`
}

// Type returns type of “C” record
func (r *CRecord) Type() string {
	return r.RecordType
}

// Parse parses the “C” record from fire ascii
func (r *CRecord) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.RecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.CRecordLayout, record)
}

// Ascii returns fire ascii of “C” record
func (r *CRecord) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.CRecordLayout)
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
func (r *CRecord) Validate() error {
	return utils.Validate(r, config.CRecordLayout, config.CRecordType)
}

// SequenceNumber returns sequence number of the record
func (r *CRecord) SequenceNumber() int {
	return r.RecordSequenceNumber
}

// SequenceNumber set sequence number of the record
func (r *CRecord) SetSequenceNumber(number int) {
	r.RecordSequenceNumber = number
}

// ControlTotal returns total of any payment amount field
func (r *CRecord) ControlTotal(index string) (int, error) {
	value, err := utils.GetField(r, "ControlTotal"+index)
	if err != nil {
		return 0, err
	}
	return int(value.Int()), nil
}

// TotalCodes returns total codes
func (r *CRecord) TotalCodes() string {
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

func (r *CRecord) ValidateRecordSequenceNumber() error {
	if r.RecordSequenceNumber < 1 {
		return utils.NewErrValidValue("sequence number")
	}
	return nil
}
