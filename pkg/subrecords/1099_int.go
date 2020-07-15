package subrecords

import (
	"reflect"
	"strings"
	"unicode/utf8"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/utils"
)

type Sub1099INT struct {
	// Enter “2” (two) to indicate notification by the IRS twice within
	// three calendar years that the payee provided an incorrect
	// name and/or TIN combination. Otherwise, enter a blank.
	SecondTinNotice string `json:"second_tin_notice"`

	// Enter the name of the foreign country or U.S. possession to
	// which the withheld foreign tax (Amount Code 6) applies.
	// Otherwise, enter blanks.
	ForeignCountry string `json:"foreign_country"`

	// Enter CUSIP Number. If the tax-exempt interest is reported
	// in the aggregate for multiple bonds or accounts, enter
	// VARIOUS. Right justify the information and fill unused
	// positions with blanks.
	CUSIPNumber string `json:"cusip_number"`

	// Enter "1" (one) if there is FATCA filing requirement.
	// Otherwise, enter a blank.
	FATCA string `json:"fatca_requirement_indicator"`

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
	CombinedFSCode string `json:"combined_federal_state_code"`
}

// Type returns type of “1099-INT” record
func (r *Sub1099INT) Type() string {
	return config.Sub1099INTType
}

// Parse parses the “1099-INT” record from fire ascii
func (r *Sub1099INT) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) < config.SubRecordLength {
		return utils.ErrSegmentLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099INTLayout, record)
}

// Ascii returns fire ascii of “1099-INT” record
func (r *Sub1099INT) Ascii() []byte {
	var buf strings.Builder
	records := config.ToSpecifications(config.Sub1099INTLayout)
	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return nil
	}

	buf.Grow(config.SubRecordLength)
	for _, spec := range records {
		value := utils.ToString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return []byte(buf.String())
}

// Validate performs some checks on the record and returns an error if not Validated
func (r *Sub1099INT) Validate() error {
	return nil
}
