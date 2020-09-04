// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package subrecords

import (
	"bytes"
	"reflect"
	"time"
	"unicode/utf8"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/utils"
)

type Sub1099B struct {
	// Enter “2” (two) to indicate notification by the IRS twice within
	// three calendar years that the payee provided an incorrect
	// name and/or TIN combination. Otherwise, enter a blank.
	SecondTinNotice string `json:"second_tin_notice"`

	// Enter the appropriate indicator from the following table, to
	// identify a Noncovered Security. If not a Noncovered Security,
	// enter a blank.
	// 1: Noncovered Security Basis not reported to the IRS
	// 2: Noncovered Security Basis reported to the IRS
	// Blank: Not a Noncovered Security
	NoncoveredSecurityIndicator string `json:"noncovered_security_indicator"`

	// Enter the appropriate indicator from the following table to
	// identify the amount reported in Amount Code 2. Otherwise,
	// enter a blank.
	// 1: Short Term
	// 2: Long Term 2
	// 3: Ordinary & Short Term
	// 4: Ordinary & Long Term
	TypeGainLossIndicator string `json:"type_gain_loss_indicator"`

	// Enter the appropriate indicator from the following table to
	// identify the amount reported in Amount Code 2. Otherwise,
	// enter a blank.
	// 1: Gross proceeds
	// 2: Gross proceeds less commissions and option premiums
	GrossProceedsIndicator string `json:"gross_proceeds_indicator"`

	// Enter blanks if this is an aggregate transaction. For broker
	// transactions, enter the trade date of the transaction. For
	// barter exchanges, enter the date when cash, property, a
	// credit, or scrip is actually or constructively received in
	// YYYYMMDD format (for example, January 5, 2019, would be
	// 20190105). Do not enter hyphens or slashes.
	DateSoldDisposed time.Time `json:"date_sold_disposed"`

	// Enter blanks if this is an aggregate transaction. Enter “0s”
	// (zeros) if the number is not available. For broker transactions
	// only, enter the CUSIP (Committee on Uniform Security
	// Identification Procedures) number of the item reported for
	// Amount Code 2 (Proceeds). Right justify the information and
	// fill unused positions with blanks.
	CUSIP string `json:"cusip_number"`

	// • For broker transactions, enter a brief description of
	//   the disposition item (e.g., 100 shares of XYZ Corp).
	// • For regulated futures and forward contracts, enter
	//   “RFC” or other appropriate description.
	// • For bartering transactions, show the services or
	//   property provided.
	// If fewer than 39 characters are required, left justify
	// information and fill unused positions with blanks.
	DescriptionProperty string `json:"description_property"`

	// Enter the date of acquisition in the format YYYYMMDD (for
	// example, January 5, 2019, would be 20190105). Do not enter
	// hyphens or slashes.
	// Enter blanks if this is an aggregate transaction.
	DateAcquired string `json:"date_acquired"`

	// Enter “1” (one) if the recipient is unable to claim a loss on
	// their tax return based on dollar amount in Amount Code 2
	// (Proceeds). Otherwise, enter a blank.
	LossNotAllowedIndicator string `json:"loss_not_allowed_indicator"`

	// Enter one of the following indicators. Otherwise, enter a blank.
	// A: Short-term transaction for which the cost or other basis is being reported to the IRS
	// B: Short-term transaction for which the cost or other basis is not being reported to the IRS
	// D: Long-term transaction for which the cost or other basis is being reported to the IRS
	// E: Long-term transaction for which the cost or other basis is not being reported to the IRS
	// X: Transaction - if you cannot determine whether the recipient should check box B or Box E on Form 8949 because the holding period is unknown
	ApplicableCheckboxForm8949 string `json:"applicable_checkbox_form8949"`

	// Enter “1” (one) if reporting proceeds from Collectibles.
	// Otherwise enter blank.
	ApplicableCheckboxCollectables string `json:"applicable_checkbox_collectables"`

	// Enter "1" (one) if there is a FATCA Filing Requirement.
	// Otherwise, enter a blank.
	FATCA string `json:"fatca_requirement_indicator"`

	// Enter a “1” (one) if reporting proceeds from QOF. Otherwise,
	// enter a blank.
	ApplicableCheckboxQOF string `json:"applicable_checkbox_qof"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for the filing requirements. If this field is
	// not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “1099-B” record
func (r *Sub1099B) Type() string {
	return config.Sub1099BType
}

// Type returns FS code of “1099-B” record
func (r *Sub1099B) FederalState() int {
	return 0
}

// Parse parses the “1099-B” record from fire ascii
func (r *Sub1099B) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099BLayout, record)
}

// Ascii returns fire ascii of “1099-B” record
func (r *Sub1099B) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099BLayout)
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
func (r *Sub1099B) Validate() error {
	return utils.Validate(r, config.Sub1099BLayout, config.Sub1099BType)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1099B) ValidateSecondTinNotice() error {
	if len(r.SecondTinNotice) > 0 &&
		r.SecondTinNotice != config.SecondTINNotice {
		return utils.NewErrValidValue("second tin notice")
	}
	return nil
}

func (r *Sub1099B) ValidateNoncoveredSecurityIndicator() error {
	if len(r.NoncoveredSecurityIndicator) > 0 &&
		(r.NoncoveredSecurityIndicator != config.GeneralOneIndicator && r.NoncoveredSecurityIndicator != config.GeneralTwoIndicator) {
		return utils.NewErrValidValue("noncovered security indicator")
	}
	return nil
}

func (r *Sub1099B) ValidateTypeGainLossIndicator() error {
	if len(r.TypeGainLossIndicator) > 0 {
		switch r.TypeGainLossIndicator {
		case "1", "2", "3", "4":
			break
		default:
			return utils.NewErrValidValue("type gain loss indicator")
		}
	}
	return nil
}

func (r *Sub1099B) ValidateGrossProceedsIndicator() error {
	if len(r.GrossProceedsIndicator) > 0 &&
		(r.GrossProceedsIndicator != config.GeneralOneIndicator && r.GrossProceedsIndicator != config.GeneralTwoIndicator) {
		return utils.NewErrValidValue("gross proceeds indicator")
	}
	return nil
}

func (r *Sub1099B) ValidateLossNotAllowedIndicator() error {
	if len(r.LossNotAllowedIndicator) > 0 &&
		r.LossNotAllowedIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("loss not allowed indicator")
	}
	return nil
}

func (r *Sub1099B) ValidateApplicableCheckboxForm8949() error {
	if len(r.ApplicableCheckboxForm8949) > 0 {
		switch r.ApplicableCheckboxForm8949 {
		case "A", "B", "D", "E", "X":
			break
		default:
			return utils.NewErrValidValue("applicable checkbox form8949")
		}
	}
	return nil
}

func (r *Sub1099B) ValidateApplicableCheckboxCollectables() error {
	if len(r.ApplicableCheckboxCollectables) > 0 &&
		r.ApplicableCheckboxCollectables != config.GeneralOneIndicator {
		return utils.NewErrValidValue("applicable checkbox collectables")
	}
	return nil
}

func (r *Sub1099B) ValidateFATCA() error {
	if len(r.FATCA) > 0 &&
		r.FATCA != config.FatcaFilingRequirementIndicator {
		return utils.NewErrValidValue("fatca filing requirement indicator")
	}
	return nil
}

func (r *Sub1099B) ValidateApplicableCheckboxQOF() error {
	if len(r.ApplicableCheckboxQOF) > 0 &&
		r.ApplicableCheckboxQOF != config.GeneralOneIndicator {
		return utils.NewErrValidValue("applicable checkbox qof")
	}
	return nil
}
