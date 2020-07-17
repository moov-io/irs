// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"github.com/moov-io/irs/pkg/subrecords"
	"reflect"
	"strings"
	"unicode/utf8"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/utils"
)

type BRecord struct {
	// Required. Enter “B.”
	RecordType string `json:"record_type" validate:"required"`

	// Required. Enter “2019.”
	// If reporting prior year data, report the year which applies (2018, 2017, etc.) and set the Prior Year Data Indicator in field position 6.
	PaymentYear int `json:"payment_year" validate:"required"`

	// Required for corrections only.
	// Indicates a corrected return. Enter the appropriate code from
	// the following table.
	// G: For a one-transaction correction or the first of a two transaction correction
	// C: For a second transaction of a two-transaction correction
	// Blank: For an original return
	// Note: C, G, and non-coded records must be reported using
	// separate Payer “A” Records.
	CorrectedReturnIndicator string `json:"corrected_return_indicator"`

	// If determinable, enter the first four characters of the last name
	// of the person whose TIN is being reported in positions 12-20
	// of the “B” Record. Otherwise, enter blanks. Last names of
	// less than four characters must be left justified and fill the
	// unused positions with blanks.
	NameControl string `json:"payees_name_control"`

	// This field is used to identify the taxpayer identification number
	// (TIN) in positions 12-20 as either an employer identification
	// number (EIN), a social security number (SSN), an individual
	// taxpayer identification number (ITIN) or an adoption taxpayer
	// identification number (ATIN). Enter the appropriate code from
	// the following table:
	// 1: EIN : A business, organization, some
	//          sole proprietors or other entity
	// 2: SSN : An individual, including some sole proprietors
	// 2: ITIN : An individual required to have a
	//           taxpayer identification number but
	//           who is not eligible to obtain an SSN
	// 2: ATIN : An adopted individual prior to the
	//           assignment of a SSN
	// Blank : N/A : If the type of TIN is not
	//               determinable, enter a blank
	TypeOfTIN string `json:"type_of_tin"`

	// Required. Enter the nine-digit taxpayer identification number
	// of the payee (SSN, ITIN, ATIN, or EIN). Do not enter hyphens
	// or alpha characters.
	// If an identification number has been applied for but not
	// received, enter blanks. All zeros, ones, twos, etc., will have
	// the effect of an incorrect TIN. If the TIN is not available, enter
	// blanks.
	TIN string `json:"payees_tin" validate:"required"`

	// Required if submitting more than one information return of the
	// same type for the same payee. Enter any number assigned by
	// the payer to the payee that can be used by the IRS to
	// distinguish between information returns. This number must be
	// unique for each information return of the same type for the
	// same payee. If a payee has more than one reporting of the
	// same document type, it is vital that each reporting have a
	// unique account number. For example, if a payer has three
	// separate pension distributions for the same payee and three
	// separate Forms 1099-R are filed; three separate unique
	// account numbers are required. A payee’s account number
	// may be given a unique sequencing number, such as 01, 02, or
	// A, B, etc., to differentiate each reported information return. Do
	// not use the payee’s TIN since this will not make each record
	// unique. This information is critical when corrections are filed.
	// This number will be provided with the backup withholding
	// notification and may be helpful in identifying the branch or
	// subsidiary reporting the transaction. The account number can
	// be any combination of alpha, numeric, or special characters. If
	// fewer than 20 characters are used, filers may either left or
	// right justify, filling the remaining positions with blanks.
	// Forms 1099-LS and 1099-SB - use this field to report
	// “Policy Number.”
	PayerAccountNumber string `json:"payers_account_number_for_payee"`

	// Enter the office code of the payer. Otherwise, enter blanks.
	// For payers with multiple locations, this field may be used to
	// identify the location of the office submitting the information
	// returns. This code will also appear on backup withholding
	// notices.
	PayerOfficeCode string `json:"payers_office_code"`

	// Required. Filers should allow for all payment amounts. For
	// those not used, enter zeros. Each payment field must contain
	// 12 numeric characters. Each payment amount must contain
	// U.S. dollars and cents. The right-most two positions represent
	// cents in the payment amount fields. Do not enter dollar signs,
	// commas, decimal points, or negative payments, except those
	// items that reflect a loss on Form 1099-B, 1099-OID, or 1099-
	// Q. Positive and negative amounts are indicated by placing a
	// “+” (plus) or “-” (minus) sign in the left-most position of the
	// payment amount field. A negative over punch in the unit’s
	// position may be used instead of a minus sign, to indicate a
	// negative amount. If a plus sign, minus sign, or negative over
	// punch is not used, the number is assumed to be positive.
	// Negative over punch cannot be used in PC created files.
	// Payment amounts must be right justified and fill unused
	// positions with zeros.
	PaymentAmount1 int `json:"payment_amount_1" validate:"required"`
	PaymentAmount2 int `json:"payment_amount_2" validate:"required"`
	PaymentAmount3 int `json:"payment_amount_3" validate:"required"`
	PaymentAmount4 int `json:"payment_amount_4" validate:"required"`
	PaymentAmount5 int `json:"payment_amount_5" validate:"required"`
	PaymentAmount6 int `json:"payment_amount_6" validate:"required"`
	PaymentAmount7 int `json:"payment_amount_7" validate:"required"`
	PaymentAmount8 int `json:"payment_amount_8" validate:"required"`
	PaymentAmount9 int `json:"payment_amount_9" validate:"required"`
	PaymentAmountA int `json:"payment_amount_A" validate:"required"`
	PaymentAmountB int `json:"payment_amount_B" validate:"required"`
	PaymentAmountC int `json:"payment_amount_C" validate:"required"`
	PaymentAmountD int `json:"payment_amount_D" validate:"required"`
	PaymentAmountE int `json:"payment_amount_E" validate:"required"`
	PaymentAmountF int `json:"payment_amount_F" validate:"required"`
	PaymentAmountG int `json:"payment_amount_G" validate:"required"`

	// If the address of the payee is in a foreign country, enter a
	// “1” (one) in this field. Otherwise, enter blank. When filers use
	// the foreign country indicator, they may use a free format for
	// the payee city, state, and ZIP Code.
	// Enter information in the following order: city, province or state,
	// postal code, and the name of the country. Do not enter
	// address information in the First or Second Payee Name Lines.
	ForeignCountryIndicator string `json:"foreign_country_indicator"`

	// Required. Enter the name of the payee (preferably last
	// name first) whose taxpayer identification number (TIN) was
	// provided in positions 12-20 of the Payee “B” Record.
	// Left justify the information and fill unused positions with
	// blanks. If more space is required for the name, use the
	// Second Payee Name Line Field. If reporting information for a
	// sole proprietor, the individual’s name must always be present
	// on the First Payee Name Line. The use of the business
	// name is optional in the Second Payee Name Line Field. End
	// the First Payee Name Line with a full word. Extraneous
	// words, titles, and special characters (that is, Mr., Mrs., Dr.,
	// period, apostrophe) should be removed from the Payee
	// Name Lines. A hyphen (-) and an ampersand (&) are the only
	// acceptable special characters for First and Second Payee
	// Name Lines.
	// Note: If a filer is required to report payments made through
	// Foreign Intermediaries and Foreign Flow-Through Entities on
	// Form 1099, see the General Instructions for Certain
	// Information Returns for reporting instructions.
	FirstPayeeNameLine string `json:"first_payee_name_line" validate:"required"`

	// If there are multiple payees (for example, partners, joint
	// owners, or spouses), use this field for those names not
	// associated with the TIN provided in positions 12-20 of the “B”
	// Record, or if not enough space was provided in the First
	// Payee Name Line, continue the name in this field. Do not
	// enter address information. It is important that filers provide as
	// much payee information to the IRS as possible to identify the
	// payee associated with the TIN. See the Note under the First
	// Payee Name Line. Left justify the information and fill unused
	// positions with blanks.
	SecondPayeeNameLine string `json:"second_payee_name_line"`

	// Required. Enter the mailing address of the payee.
	// The street address should include number, street, apartment
	// or suite number, or P.O. Box if mail is not delivered to a
	// street address. Left justify the information and fill unused
	// positions with blanks.
	// Do not enter data other than the payee’s mailing address.
	PayeeMailingAddress string `json:"payee_mailing_address" validate:"required"`

	// Required. Enter the city, town or post office. Enter APO or
	// FPO if applicable. Do not enter state and ZIP Code
	// information in this field. Left justify the information and fill
	// unused positions with blanks.
	PayeeCity string `json:"payee_city" validate:"required"`

	// Required. Enter the valid U.S. Postal Service state
	// abbreviations for states or the appropriate postal identifier
	// (AA, AE, or AP).
	PayeeState string `json:"payee_state" validate:"required"`

	// Required. Enter the valid ZIP Code (nine-digit or five-digit)
	// assigned by the U.S. Postal Service.
	// For foreign countries, alpha characters are acceptable as
	// long as the filer has entered a “1” (one) in the Foreign
	// Country Indicator, located in position 247 of the “B” Record. If
	// only the first five-digits are known, left justify the information
	// and fill the unused positions with blanks.
	PayeeZipCode string `json:"payee_zip_code" validate:"required"`

	// Required. Enter the number of the record as it appears
	//within the file. The record sequence number for the “T”
	//Record will always be one (1), since it is the first record on
	//the file and the file can have only one “T” Record in a file.
	//Each record, thereafter, must be increased by one in
	//ascending numerical sequence, that is, 2, 3, 4, etc. Right
	//justify numbers with leading zeros in the field. For example,
	//the “T” Record sequence number would appear as
	//“00000001” in the field, the first “A” Record would be
	//“00000002,” the first “B” Record, “00000003,” the second “B”
	//Record, “00000004”, and so on until the final record of the
	//file, the “F” Record.
	RecordSequenceNumber int `json:"record_sequence_number" validate:"required"`

	typeOfReturn string
	extRecord    subrecords.SubRecord
}

// Type returns type of “B” record
func (r *BRecord) Type() string {
	return config.BRecordType
}

// Parse parses the “B” record from fire ascii
func (r *BRecord) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.RecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.BRecordLayout, record)
}

// Ascii returns fire ascii of “B” record
func (r *BRecord) Ascii() []byte {
	var buf strings.Builder
	records := config.ToSpecifications(config.BRecordLayout)
	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return nil
	}

	buf.Grow(config.RecordLength)
	for _, spec := range records {
		value := utils.ToString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return []byte(buf.String())
}

// Validate performs some checks on the record and returns an error if not Validated
func (r *BRecord) Validate() error {
	return utils.Validate(r, config.BRecordLayout)
}

// SequenceNumber returns sequence number of the record
func (r *BRecord) SequenceNumber() int {
	return r.RecordSequenceNumber
}

// SequenceNumber set sequence number of the record
func (r *BRecord) SetSequenceNumber(number int) {
	r.RecordSequenceNumber = number
}

// SetTypeOfReturn set type of return of the record
func (r *BRecord) SetTypeOfReturn(typeOfReturn string) {
	r.typeOfReturn = typeOfReturn
}

// SetTypeOfReturn returns type of return of the record
func (r *BRecord) TypeOfReturn() string {
	return r.typeOfReturn
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *BRecord) ValidateSequenceNumber() error {
	if r.RecordSequenceNumber < 1 {
		return utils.NewErrValidValue("sequence number")
	}
	return nil
}
