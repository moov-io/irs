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

type Sub1099K struct {
	// Enter “2” (two) to indicate notification by the IRS twice within
	// three calendar years that the payee provided an incorrect
	// name and/or TIN combination. Otherwise, enter a blank.
	SecondTinNotice string `json:"second_tin_notice"`

	// Required. Enter the appropriate indicator from the following
	// table.
	// 1: Payment Settlement Entity (PSE)
	// 2: Electronic Payment Facilitator (EPF)/Other third party
	TypeFilerIndicator string `json:"type_filer_indicator"`

	// Required. Enter the appropriate indicator from the following
	// table.
	// 1: Payment Card Payment
	// 2: EThird Party Network Payment
	TypePaymentIndicator string `json:"type_payment_indicator"`

	// Required. Enter the number of payment transactions. Do not
	// include refund transactions.
	// Right justify the information and fill unused positions with
	// zeros.
	NumberPaymentTransactions int `json:"number_payment_transactions"`

	// Enter the payment settlement entity’s name and phone
	// number if different from the filer's name. Otherwise, enter
	// blanks. Left justify the information, and fill unused positions
	// with blanks.
	PaymentSettlementNamePhoneNumber string `json:"payment_settlement_name_phone_number"`

	// Required. Enter the Merchant Category Code (MCC). All
	// MCCs must contain four numeric characters. If no code is
	// provided, fill unused positions with zeros.
	MerchantCategoryCode int `json:"merchant_category_code"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for filing requirements. You may enter
	// your routing and transit number (RTN) here. If this field is not
	// used, enter blanks.
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

// Type returns type of “1099-K” record
func (r *Sub1099K) Type() string {
	return config.Sub1099KType
}

// Type returns FS code of “1099-K” record
func (r *Sub1099K) FederalState() int {
	return r.CombinedFSCode
}

// Parse parses the “1099-K” record from fire ascii
func (r *Sub1099K) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099KLayout, record)
}

// Ascii returns fire ascii of “1099-K” record
func (r *Sub1099K) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099KLayout)
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
func (r *Sub1099K) Validate() error {
	return utils.Validate(r, config.Sub1099KLayout, config.Sub1099KType)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1099K) ValidateSecondTinNotice() error {
	if len(r.SecondTinNotice) > 0 &&
		r.SecondTinNotice != config.SecondTINNotice {
		return utils.NewErrValidValue("second tin notice")
	}
	return nil
}

func (r *Sub1099K) ValidateTypeFilerIndicator() error {
	if r.TypeFilerIndicator != config.GeneralOneIndicator &&
		r.TypeFilerIndicator != config.GeneralTwoIndicator {
		return utils.NewErrValidValue("type filer indicator")
	}
	return nil
}

func (r *Sub1099K) ValidateTypePaymentIndicator() error {
	if r.TypePaymentIndicator != config.GeneralOneIndicator &&
		r.TypePaymentIndicator != config.GeneralTwoIndicator {
		return utils.NewErrValidValue("type payment indicator")
	}
	return nil
}

func (r *Sub1099K) ValidateCombinedFSCode() error {
	return utils.ValidateCombinedFSCode(r.CombinedFSCode)
}
