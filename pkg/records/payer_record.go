// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"bytes"
	"reflect"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/utils"
)

type ARecord struct {
	// Required. Enter “A.”
	RecordType string `json:"record_type" validate:"required"`

	// Required. Enter “2019.”
	// If reporting prior year data, report the year which applies (2018, 2017, etc.) and set the Prior Year Data Indicator in field position 6.
	PaymentYear int `json:"payment_year" validate:"required"`

	// Required for CF/SF.
	// Enter “1” (one) if approved and submitting information as part
	// of the CF/SF Program or if submitting a test file in order to
	// obtain approval for the CF/SF Program. Otherwise, enter a
	// blank.
	// Note 1: If the Payer “A” Record is coded for CF/SF, there
	// must be coding in the Payee “B” Records and the State Totals
	// “K” Records.
	// Note 2: If “1” (one) is entered in this field position, be sure to
	// code the Payee “B” Records with the appropriate state code.
	// Refer to Part A. Sec. 12, Table 1, Participating States and
	// Codes, for further information.
	CombinedFSFilingProgram string `json:"combined_fs_filing_program"`

	// Required. Enter the valid nine-digit taxpayer identification
	// number assigned to the payer. Do not enter blanks, hyphens,
	// or alpha characters. Filling the field with all zeros, ones, twos,
	// etc., will result in an incorrect TIN.
	// Note: For foreign entities that are not required to have a TIN,
	// this field must be blank; however, the Foreign Entity Indicator,
	// position 52 of the “A” Record, must be set to one (1).
	TIN string `json:"payer_tin" validate:"required"`

	// Enter the four characters of the name control or enter blanks.
	PayerNameControl string `json:"payer_name_control"`

	// Enter “1” (one) if this is the last year this payer name and TIN
	// will file information returns electronically or on paper.
	// Otherwise, enter a blank.
	LastFilingIndicator string `json:"last_filing_indicator"`

	// Required. Enter the appropriate code. Left justify and fill
	// unused positions with blanks.
	TypeOfReturn string `json:"type_of_return" validate:"required"`

	// Required. Enter the appropriate amount code(s) for the type
	// of return being reported. In most cases, the box numbers on
	// paper information returns correspond with the amount codes
	// used to file electronically. However, if discrepancies occur,
	// Publication 1220 governs for filing electronically. Enter the
	// amount codes in ascending sequence; numeric characters
	// followed by alphas. Left justify the information and fill unused
	// positions with blanks.
	// Note: A type of return and an amount code must be present
	// in every Payer “A” Record even if no money amounts are
	// being reported. For a detailed explanation of the information
	// to be reported in each amount code, refer to the appropriate
	// paper instructions for each form.
	AmountCodes string `json:"amount_codes" validate:"required"`

	// Enter “1” (one) if the transmitter is a foreign entity. If the transmitter is not a foreign entity, enter a blank.
	ForeignEntityIndicator string `json:"foreign_entity_indicator"`

	// Required. Enter the name of the payer whose TIN appears in
	// positions 12-20 of the “A” Record. (The transfer agent’s name
	// is entered in the Second Payer Name Line Field, if
	// applicable). Left justify information and fill unused positions
	// with blanks. Delete extraneous information.
	FirstPayerNameLine string `json:"first_payer_name" validate:"required"`

	// If position 133 Transfer (or Paying) Agent Indicator contains a
	// “1” (one), this field must contain the name of the transfer or
	// paying agent.
	// If position 133 contains a “0” (zero), this field may contain
	// either a continuation of the First Payer Name Line or blanks.
	// Left justify the information. Fill unused positions with blanks
	SecondPayerNameLine string `json:"second_payer_name"`

	// Required. Enter the appropriate numeric code from the table below
	// 1: The entity in the Second Payer Name Line Field is the transfer (or paying) agent.
	// 0: The entity shown is not the transfer (or paying) agent (that is, the Second Payer Name Line Field either contains
	//    a continuation of the First Payer Name Line Field or blanks).
	TransferAgentIndicator string `json:"transfer_agent_control" validate:"required"`

	// Required. If position 133 Transfer Agent Indicator is “1” (one),
	// enter the shipping address of the transfer or paying agent.
	// Otherwise, enter the actual shipping address of the payer. The
	// street address includes street number, apartment or suite
	// number, or P.O. Box address if mail is not delivered to a street
	// address. Left justify the information and fill unused positions
	// with blanks.
	// For U.S. addresses, the payer city, state, and ZIP Code must
	// be reported as 40-, 2-, and 9-position fields, respectively.
	// Filers must adhere to the correct format for the payer city,
	// state, and ZIP Code.
	// For foreign addresses, filers may use the payer city, state, and
	// ZIP Code as a continuous 51-position field. Enter information
	// in the following order: city, province or state, postal code, and
	// the name of the country. When reporting a foreign address,
	// the Foreign Entity Indicator in position 52 must contain a
	// "1" (one).
	PayerShippingAddress string `json:"payer_shipping_address" validate:"required"`

	// Required. If the Transfer Agent Indicator in position 133 is a
	// “1” (one), enter the city, town, or post office of the transfer
	// agent. Otherwise, enter payer’s city, town, or post office city.
	// Do not enter state and ZIP Code information in this field. Left
	// justify the information and fill unused positions with blanks.
	PayerCity string `json:"payer_city" validate:"required"`

	// Required. Enter the valid U.S. Postal Service state abbreviation.
	PayerState string `json:"payer_state" validate:"required"`

	// Required. Enter the valid nine-digit ZIP Code assigned by the
	// U.S. Postal Service. If only the first five digits are known, left
	// justify the information and fill unused positions with blanks. For
	// foreign countries, alpha characters are acceptable as long as
	// the filer has entered a “1” (one) in “A” Record, field position 52
	// Foreign Entity Indicator.
	PayerZipCode string `json:"payer_zip_code" validate:"required"`

	// Enter the payer’s telephone number and extension. Omit
	// hyphens. Left justify the information and fill unused positions
	// with blanks.
	PayerTelephoneNumber string `json:"payer_telephone_number_and_ext"`

	// Required. Enter the number of the record as it appears within
	// the file. The record sequence number for the “T” Record will
	// always be “1” (one), since it is the first record on the file and
	// the file can have only one “T” Record. Each record thereafter
	// must be increased by one in ascending numerical sequence,
	// that is, 2, 3, 4, etc. Right justify numbers with leading zeros in
	// the field. For example, the “T” Record sequence number
	// would appear as “00000001” in the field, the first “A” Record
	// would be “00000002,” the first “B” Record, “00000003,” the
	// second “B” Record, “00000004” and so on until the final record
	// of the file, the “F” Record.
	RecordSequenceNumber int `json:"record_sequence_number" validate:"required"`
}

// Type returns type of “A” record
func (r *ARecord) Type() string {
	return config.ARecordType
}

// Parse parses the “A” record from fire ascii
func (r *ARecord) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.RecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.ARecordLayout, record)
}

// Ascii returns fire ascii of “A” record
func (r *ARecord) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.ARecordLayout)
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
func (r *ARecord) Validate() error {
	return utils.Validate(r, config.ARecordLayout)
}

// SequenceNumber returns sequence number of the record
func (r *ARecord) SequenceNumber() int {
	return r.RecordSequenceNumber
}

// SequenceNumber set sequence number of the record
func (r *ARecord) SetSequenceNumber(number int) {
	r.RecordSequenceNumber = number
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *ARecord) ValidateRecordSequenceNumber() error {
	if r.RecordSequenceNumber < 1 {
		return utils.NewErrValidValue("sequence number")
	}
	return nil
}

func (r *ARecord) ValidateCombinedFSFilingProgram() error {
	if r.CombinedFSFilingProgram == config.FSFilingProgramApproved || len(r.CombinedFSFilingProgram) == 0 {
		return nil
	}
	return utils.NewErrValidValue("combined federal filing program")
}

func (r *ARecord) ValidateLastFilingIndicator() error {
	if r.LastFilingIndicator == config.LastFilingIndicator || len(r.LastFilingIndicator) == 0 {
		return nil
	}
	return utils.NewErrValidValue("last filing indicator")
}

func (r *ARecord) ValidateTypeOfReturn() error {
	if _, ok := config.TypeOfReturns[r.TypeOfReturn]; ok {
		return nil
	}
	return utils.NewErrValidValue("type of return")
}

func (r *ARecord) ValidateForeignEntityIndicator() error {
	if r.ForeignEntityIndicator == config.ForeignEntityIndicator || len(r.ForeignEntityIndicator) == 0 {
		return nil
	}
	return utils.NewErrValidValue("foreign entity indicator")
}

func (r *ARecord) ValidateTransferAgentIndicator() error {
	if r.TransferAgentIndicator == config.TransferAgentIndicator || r.TransferAgentIndicator == config.NotTransferAgentIndicator {
		return nil
	}
	return utils.NewErrValidValue("transfer agent indicator")
}

func (r *ARecord) ValidatePayerState() error {
	if _, ok := config.StateAbbreviationCodes[r.PayerState]; ok {
		return nil
	}
	return utils.NewErrValidValue("payer state")
}

func (r *ARecord) ValidateAmountCodes() error {
	returnType, exist := config.TypeOfReturns[r.TypeOfReturn]
	if !exist {
		return utils.NewErrValidValue("type of return")
	}

	codeMap, exist := config.AmountCodes[returnType]
	if !exist {
		return utils.NewErrValidValue("amount codes")
	}
	if !checkAvailableCodes(r.AmountCodes, codeMap) {
		return utils.NewErrValidValue("amount codes")
	}

	return nil
}

func checkAvailableCodes(codes string, codeMap map[string]string) bool {
	codes = strings.TrimRight(codes, config.BlankString)
	codeList := strings.Split(codes, "")
	sort.Strings(codeList)
	if strings.Join(codeList, "") != codes {
		return false
	}

	repeated := map[string]int{}
	for i := 0; i < len(codeList); i++ {
		repeated[codeList[i]]++
	}

	for code, v := range repeated {
		if v > 1 {
			return false
		}
		if _, ok := codeMap[code]; !ok {
			return false
		}
	}

	return true
}
