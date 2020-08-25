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

type Sub1099LTC struct {
	// Enter the appropriate indicator from the following table.
	// Otherwise, enter blanks.
	// 1: Per diem
	// 2: Reimbursed amount
	TypePaymentIndicator string `json:"type_payment_indicator"`

	// Required. Enter the social security number of the insured.
	SocialSecurityNumberInsured string `json:"social_security_number_insured"`

	// Required. Enter the name of the insured.
	NameInsured string `json:"name_insured"`

	// Required. Enter the address of the insured. The street
	// address should include number, street, apartment or suite
	// number (or P.O. Box if mail is not delivered to street
	// address). Do not input any data other than the payee’s
	// address. Left justify the information and fill unused positions
	// with blanks.
	// For U.S. addresses, the payee city, state, and ZIP Code
	// must be reported as a 40-, 2-, and 9-position field,
	// respectively. Filers must adhere to the correct format for the
	// insured’s city, state, and ZIP Code.
	// For foreign addresses, filers may use the insured’s city,
	// state, and ZIP Code as a continuous 51-position field. Enter
	// information in the following order: city, province or state,
	// postal code, and the name of the country. When reporting a
	// foreign address, the Foreign Country Indicator in position 247
	// must contain a “1” (one).
	AddressInsured string `json:"address_insured"`

	// Required. Enter the city, town, or post office. Left justify the
	// information and fill unused positions with blanks. Enter APO
	// or FPO, if applicable. Do not enter state and ZIP Code
	// information in this field. Left justify the information and fill
	// unused positions with blanks.
	CityInsured string `json:"city_insured"`

	// Required. Enter the valid U.S. Postal Service state
	// abbreviations for states or the appropriate postal identifier
	// (AA, AE, or AP). Refer to Part A. Sec. 13, Table 2, State &
	// U.S. Territory Abbreviations
	StateInsured string `json:"state_insured"`

	// Required. Enter the valid nine-digit ZIP Code assigned by
	// the U.S. Postal Service. If only the first five-digits are known,
	// left justify the information and fill the unused positions with
	// blanks. For foreign countries, alpha characters are
	// acceptable as long as the filer has entered a “1” (one) in the
	// Foreign Country Indicator, located in position 247 of the “B”
	// Record.
	ZipCodeInsured string `json:"zip_code_insured"`

	// Enter the appropriate code from the table below to indicate
	// the status of the illness of the insured. Otherwise, enter
	// blank.
	// 1: Chronically ill
	// 2: Terminally ill
	StatusIllnessIndicator string `json:"status_illness_indicator"`

	// Enter the latest date of a doctor’s certification of the status of
	// the insured’s illness in YYYYMMDD format (for example,
	// January 5, 2019, would be 20190105). Do not enter hyphens
	// or slashes.
	DateCertified string `json:"date_certified"`

	// Enter “1” (one) if benefits were from a qualified long-term
	// care insurance contract. Otherwise, enter a blank.
	QualifiedContractIndicator string `json:"qualified_contract_indicator"`

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
}

// Type returns type of “1099-LTC” record
func (r *Sub1099LTC) Type() string {
	return config.Sub1099LtcType
}

// Parse parses the “1099-LTC” record from fire ascii
func (r *Sub1099LTC) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099LTCLayout, record)
}

// Ascii returns fire ascii of “1099-LTC” record
func (r *Sub1099LTC) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099LTCLayout)
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
func (r *Sub1099LTC) Validate() error {
	return utils.Validate(r, config.Sub1099LTCLayout)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1099LTC) ValidateTypePaymentIndicator() error {
	if len(r.TypePaymentIndicator) > 0 &&
		(r.TypePaymentIndicator != config.GeneralOneIndicator && r.TypePaymentIndicator != config.GeneralTwoIndicator) {
		return utils.NewErrValidValue("type payment indicator")
	}
	return nil
}

func (r *Sub1099LTC) ValidateStatusIllnessIndicator() error {
	if len(r.StatusIllnessIndicator) > 0 &&
		(r.StatusIllnessIndicator != config.GeneralOneIndicator && r.StatusIllnessIndicator != config.GeneralTwoIndicator) {
		return utils.NewErrValidValue("status illness indicator")
	}
	return nil
}

func (r *Sub1099LTC) ValidateQualifiedContractIndicator() error {
	if len(r.QualifiedContractIndicator) > 0 &&
		r.QualifiedContractIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("qualified contract indicator")
	}
	return nil
}

func (r *Sub1099LTC) ValidateStateInsured() error {
	if _, ok := config.StateAbbreviationCodes[r.StateInsured]; !ok {
		return utils.NewErrValidValue("state insured")
	}
	return nil
}
