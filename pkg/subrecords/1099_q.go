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

type Sub1099Q struct {
	// Required. Enter “1” (one) if reporting a trustee to trustee
	// transfer. Otherwise, enter a blank.
	TrusteeTransferIndicator string `json:"trustee_transfer_indicator"`

	// Required. Enter the appropriate code from the table below to
	// indicate the type of tuition payment. Otherwise, enter a
	// blank.
	// 1: Private program payment
	// 2: State program payment
	// 3: Coverdell ESA contribution
	TypeTuitionPayment string `json:"type_tuition_payment"`

	// Required. Enter “1” (one) if the recipient is not the
	// designated beneficiary. Otherwise, enter a blank.
	DesignatedBeneficiary string `json:"designated_beneficiary"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for filing requirements. You may enter
	// your routing and transit number (RTN) here. If this field is not
	// used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “1099-Q” record
func (r *Sub1099Q) Type() string {
	return config.Sub1099QType
}

// Type returns FS code of “1099-Q” record
func (r *Sub1099Q) FederalState() int {
	return 0
}

// Parse parses the “1099-Q” record from fire ascii
func (r *Sub1099Q) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099QLayout, record)
}

// Ascii returns fire ascii of “1099-Q” record
func (r *Sub1099Q) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099QLayout)
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
func (r *Sub1099Q) Validate() error {
	return utils.Validate(r, config.Sub1099QLayout, config.Sub1099QType)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1099Q) ValidateTrusteeTransferIndicator() error {
	if len(r.TrusteeTransferIndicator) > 0 &&
		r.TrusteeTransferIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("trustee transfer indicator")
	}
	return nil
}

func (r *Sub1099Q) ValidateDesignatedBeneficiary() error {
	if len(r.DesignatedBeneficiary) > 0 &&
		(r.DesignatedBeneficiary != config.GeneralOneIndicator && r.DesignatedBeneficiary != config.GeneralTwoIndicator) {
		return utils.NewErrValidValue("designated beneficiary")
	}
	return nil
}

func (r *Sub1099Q) ValidateTypeTuitionPayment() error {
	if len(r.TypeTuitionPayment) > 0 {
		switch r.TypeTuitionPayment {
		case "1", "2", "3":
			return nil
		default:
			return utils.NewErrValidValue("type tuition payment")
		}
	}
	return nil
}
