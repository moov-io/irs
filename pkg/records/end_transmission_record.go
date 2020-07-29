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

type FRecord struct {
	// Required. Enter “F.”
	RecordType string `json:"record_type" validate:"required"`

	// Enter zeros.
	Zero int `json:"zero"`

	// Enter the total number of Payer “A” Records in the entire file.
	// Right justify the information and fill unused positions with
	// zeros or enter all zeros.
	NumberPayerRecords int `json:"number_of_payer_records"`

	// If this total was entered in the “T” Record, this field may be
	// blank filled. Enter the total number of Payee “B” Records
	// reported in the file. Right justify the information and fill
	// unused positions with zeros.
	TotalNumberPayees int `json:"total_number_of_payees"`

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

// Type returns type of “F” record
func (r *FRecord) Type() string {
	return config.FRecordType
}

// Parse parses the “F” record from fire ascii
func (r *FRecord) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.RecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.FRecordLayout, record)
}

// Ascii returns fire ascii of “F” record
func (r *FRecord) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.FRecordLayout)
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
func (r *FRecord) Validate() error {
	return utils.Validate(r, config.FRecordLayout)
}

// SequenceNumber returns sequence number of the record
func (r *FRecord) SequenceNumber() int {
	return r.RecordSequenceNumber
}

// SequenceNumber set sequence number of the record
func (r *FRecord) SetSequenceNumber(number int) {
	r.RecordSequenceNumber = number
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *FRecord) ValidateRecordSequenceNumber() error {
	if r.RecordSequenceNumber < 1 {
		return utils.NewErrValidValue("sequence number")
	}
	return nil
}
