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

type Sub5498 struct {
	// Enter “1” (one) if reporting a rollover (Amount Code 2) or Fair
	// Market Value (Amount Code 5) for an IRA. Otherwise, enter
	// a blank.
	IRAIndicator string `json:"ira_indicator"`

	// Enter “1” (one) if reporting a rollover (Amount Code 2) or Fair
	// Market Value (Amount Code 5) for a SEP. Otherwise, enter
	// a blank.
	SEPIndicator string `json:"sep_indicator"`

	// Enter “1” (one) if reporting a rollover (Amount Code 2) or Fair
	// Market Value (Amount Code 5) for a SIMPLE. Otherwise,
	// enter a blank.
	SIMPLEIndicator string `json:"simple_indicator"`

	// Enter “1” (one) if reporting a rollover (Amount Code 2) or Fair
	// Market Value (Amount Code 5) for a Roth IRA. Otherwise,
	// enter a blank.
	RothIRAIndicator string `json:"roth_ira_indicator"`

	// Enter “1” (one) if reporting RMD for 2020. Otherwise, enter a
	// blank.
	RMDIndicator string `json:"rmd_indicator"`

	// Required. Enter the date the option was granted in
	// YYYYMMDD format (for example, January 5, 2019, would be
	// 20190105).
	YearPostponedContribution int `json:"year_postponed_contribution"`

	// Required, if applicable. Enter the code from the table below.
	// Right justify. Otherwise, enter blanks.
	// FD: Federally Designated Disaster Area
	// PL: Public Law
	// EO: Executive Order
	// PO: Rollovers of qualified plan loan offset amounts
	// SC: For participants who have certified that
	//     the rollover contribution is late because
	//     of an error on the part of a financial
	//     institution, death, disability,
	//     hospitalization, incarceration,
	//     restrictions imposed by a foreign
	//     country, postal error, or other
	//     circumstance listed in Section
	//     3.02(2) of Rev. Proc. 2016-47 or other
	//     event beyond the reasonable control of
	//     the participant.
	PostponedContributionCode string `json:"postponed_contribution_code"`

	// Required, if applicable. Enter the federally declared disaster
	// area, public law number or executive order number under
	// which the postponed contribution is being issued.
	// Right justify. Otherwise, enter blanks.
	PostponedContributionReason string `json:"postponed_contribution_reason"`

	// Required. Enter the two-character alpha Repayment Code.
	// Right justify. Otherwise, enter blanks.
	// QR: Qualified Reservist Distribution
	// DD: Federally Designated Disaster Distribution
	RepaymentCode string `json:"repayment_code"`

	// Enter the date by which the RMD amount must be distributed
	// to avoid the 50% excise tax. Format the date as
	// YYYYMMDD (for example, January 5, 2019, would be
	// 20190105). Otherwise, enter blanks.
	RMDDate time.Time `json:"rmd_date"`

	// Equal to one alpha character or two alpha characters or
	// blank. Valid characters are:
	// • Two-character combinations can consist of A, B, C,
	//   D, E, F, and G.
	// • Valid character H cannot be present with any other
	//  characters.
	Codes string `json:"codes"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for filing requirements.
	// If this field is not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`

	// Enter the valid CF/SF code if this payee record is to be
	// forwarded to a state agency as part of the CF/SF Program.
	CombinedFSCode int `json:"combined_federal_state_code"`
}

// Type returns type of “5498” record
func (r *Sub5498) Type() string {
	return config.Sub5498Type
}

// Parse parses the “5498” record from fire ascii
func (r *Sub5498) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub5498Layout, record)
}

// Ascii returns fire ascii of “5498” record
func (r *Sub5498) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub5498Layout)
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
func (r *Sub5498) Validate() error {
	return utils.Validate(r, config.Sub5498Layout)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub5498) ValidateIRAIndicator() error {
	if len(r.IRAIndicator) > 0 &&
		r.IRAIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("ira indicator")
	}
	return nil
}

func (r *Sub5498) ValidateSEPIndicator() error {
	if len(r.SEPIndicator) > 0 &&
		r.SEPIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("sep indicator")
	}
	return nil
}

func (r *Sub5498) ValidateSIMPLEIndicator() error {
	if len(r.SIMPLEIndicator) > 0 &&
		r.SIMPLEIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("simple indicator")
	}
	return nil
}

func (r *Sub5498) ValidateRothIRAIndicator() error {
	if len(r.RothIRAIndicator) > 0 &&
		r.RothIRAIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("roth ira indicator")
	}
	return nil
}

func (r *Sub5498) ValidateRMDIndicator() error {
	if len(r.RMDIndicator) > 0 &&
		r.RMDIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("rmd indicator")
	}
	return nil
}

func (r *Sub5498) ValidatePostponedContributionCode() error {
	if len(r.PostponedContributionCode) > 0 {
		switch r.PostponedContributionCode {
		case "FD", "PL", "EO", "PO", "SC":
			return nil
		default:
			return utils.NewErrValidValue("postponed contribution code")
		}
	}
	return nil
}

func (r *Sub5498) ValidateRepaymentCode() error {
	if len(r.RepaymentCode) > 0 {
		switch r.RepaymentCode {
		case "QR", "DD":
			return nil
		default:
			return utils.NewErrValidValue("repayment code")
		}
	}
	return nil
}

func (r *Sub5498) ValidateCodes() error {
	if len(r.Codes) > 0 {
		lowCode := 'A'
		highCode := 'H'
		if len(r.Codes) > 1 {
			highCode = 'G'
		}
		for _, letter := range r.Codes {
			if letter >= lowCode && letter <= highCode {
				return nil
			} else {
				return utils.NewErrValidValue("repayment code")
			}
		}
	}
	return nil
}

func (r *Sub5498) ValidateCombinedFSCode() error {
	if _, ok := config.ParticipateStateCodes[r.CombinedFSCode]; !ok {
		return utils.NewErrValidValue("combined federal state code")
	}
	return nil
}
