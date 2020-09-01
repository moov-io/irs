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

type Sub1099G struct {
	// Enter “2” (two) to indicate notification by the IRS twice within
	// three calendar years that the payee provided an incorrect
	// name and/or TIN combination. Otherwise, enter a blank.
	SecondTinNotice string `json:"second_tin_notice"`

	// Enter “1” (one) to indicate the state or local income tax
	// refund, credit, or offset (Amount Code 2) is attributable to
	// income tax that applies exclusively to income from a trade or
	// business.
	// 1: Income tax refund applies exclusively to a trade or business
	// Blank: Income tax refund is a general tax refund
	TradeBusinessIndicator string `json:"trade_business_indicator"`

	// Enter the tax year for which the refund, credit, or offset
	// (Amount Code 2) was issued. The tax year must reflect the
	// tax year for which the refund was made, not the tax year of
	// Form 1099-G. The tax year must be in four-position format of
	// YYYY (for example, 2015). The valid range of years for the
	// refund is 2009 through 2018.
	// Note: This data is not considered prior year data since it is
	// required to be reported in the current tax year. Do NOT enter
	// “P” in the field position 6 of Transmitter “T” Record.
	TaxYearRefund int `json:"tax_tear_refund"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for the filing requirements. If this field is
	// not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`

	// State income tax withheld is for the convenience of the filers.
	// This information does not need to be reported to the IRS. If
	// not reporting state tax withheld, this field may be used as a
	// continuation of the Special Data Entries Field. The payment
	// amount must be right justified and unused positions
	// zero-filled.
	StateIncomeTaxWithheld int `json:"state_income_tax_withheld"`

	// Local income tax withheld is for the convenience of the filers.
	// This information does not need to be reported to the IRS. If
	// not reporting local tax withheld, this field may be used as a
	// continuation of the Special Data Entries Field. The payment
	// amount must be right justified and unused positions
	// zero-filled.
	LocalIncomeTaxWithheld int `json:"local_income_tax_withheld"`

	// Enter the valid CF/SF code if this payee record is to be
	// forwarded to a state agency as part of the CF/SF Program.
	CombinedFSCode int `json:"combined_federal_state_code"`
}

// Type returns type of “1099-G” record
func (r *Sub1099G) Type() string {
	return config.Sub1099GType
}

// Parse parses the “1099-G” record from fire ascii
func (r *Sub1099G) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099GLayout, record)
}

// Ascii returns fire ascii of “1099-G” record
func (r *Sub1099G) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099GLayout)
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
func (r *Sub1099G) Validate() error {
	return utils.Validate(r, config.Sub1099GLayout, config.Sub1099GType)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1099G) ValidateSecondTinNotice() error {
	if len(r.SecondTinNotice) > 0 &&
		r.SecondTinNotice != config.SecondTINNotice {
		return utils.NewErrValidValue("second tin notice")
	}
	return nil
}

func (r *Sub1099G) ValidateTradeBusinessIndicator() error {
	if len(r.TradeBusinessIndicator) > 0 &&
		r.TradeBusinessIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("trade business indicator")
	}
	return nil
}

func (r *Sub1099G) ValidateCombinedFSCode() error {
	if _, ok := config.ParticipateStateCodes[r.CombinedFSCode]; !ok {
		return utils.NewErrValidValue("combined federal state code")
	}
	return nil
}

func (r *Sub1099G) ValidateTaxYearRefund() error {
	if r.TaxYearRefund < 2009 || r.TaxYearRefund > 2018 {
		return utils.NewErrValidValue("tax tear refund")
	}
	return nil
}
