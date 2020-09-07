// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package subrecords

import (
	"bytes"
	"encoding/json"
	"reflect"
	"time"
	"unicode/utf8"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/utils"
)

type Sub1099R struct {
	// Required. Enter at least one distribution code from the table
	// below. More than one code may apply. If only one code is
	// necessary, it must be entered in position 545 and position
	// 546 will be blank. When using Code P for an IRA distribution
	// under Section 408(d)(4) of the Internal Revenue Code, the
	// filer may also enter Code 1, 2, 4, B or J, if applicable. Only
	// three numeric combinations are acceptable; Codes 8 and 1,
	// 8 and 2, and 8 and 4, on one return. These three
	// combinations can be used only if both codes apply to the
	// distribution being reported. If more than one numeric code is
	// applicable to different parts of a distribution, report two
	// separate “B” Records.
	// • Distribution Codes 5, 9, E, F, N, Q, R, S and T cannot be
	//   used with any other codes.
	// • Distribution Code C can be a stand alone or combined
	//   with Distribution Code D only.
	// • Distribution Code G may be used with Distribution Code
	//   4 only if applicable.
	// • Distribution Code K is valid with Distribution Codes 1, 2,
	//   4, 7, 8, or G.
	// • Distribution Code M can be a stand alone or combined
	//   with Distribution Codes 1, 2, 4, 7, or B.
	// 1: *Early distribution, no known exception (in most cases, under age 59½)
	// 2: *Early distribution, exception applies (under age 59½)
	// 3: *Disability
	// 4: *Death
	// 5: *Prohibited transaction
	// 6: Section 1035 exchange (a taxfree exchange of life insurance, annuity, qualified long-term care insurance, or endowment contracts)
	// 7: *Normal distribution
	// 8: *Excess contributions plus earnings/excess deferrals (and/or earnings) taxable in 2019
	// 9: Cost of current life insurance protection (premiums paid by a trustee or custodian for current insurance protection)
	// A: May be eligible for 10-year tax option
	// B: Designated Roth account distribution
	// C: Reportable Death Benefits under Section 6050Y(c)
	// D: Annuity payments from nonqualified annuity payments and distributions from life insurance contracts that may be subject to tax under Section 1411
	// E: Distribution under Employee Plans Compliance Resolution System (EPCRS)
	// F: Charitable gift annuity
	// G: Direct rollover and rollover contribution
	// H: Direct rollover of distribution from a designated Roth account to a Roth IRA
	// J: Early distribution from a Roth IRA (This code may be used with a Code 8 or P)
	// K: Distribution of IRA assets not having a readily available FMV
	// L: Loans treated as deemed distributions under Section 72(p)
	// M: Qualified Plan Loan Offsets
	// N: Recharacterized IRA contribution made for 2019
	// P: *Excess contributions plus earnings/excess deferrals taxable for 2018
	// Q: Qualified distribution from a Roth IRA.
	// R: Recharacterized IRA contribution made for 2018
	// S: *Early distribution from a SIMPLE IRA in first 2 years no known exceptions
	// T: Roth IRA distribution exception applies because participant has reached 59½, died or is disabled, but it is unknown if the 5-year  period has been met
	// U: Distribution from ESOP under Section 404(k)
	// W: Charges or payments for purchasing qualified long-term care insurance contracts under combined arrangements
	DistributionCode string `json:"distribution_code"`

	// Enter “1” (one) only if the taxable amount of the payment
	// entered for Payment Amount Field 1 (Gross distribution) of
	// the “B” Record cannot be computed. Otherwise, enter a
	// blank. (If the Taxable Amount Not Determined Indicator is
	// used, enter “0s” [zeros] in Payment Amount Field 2 of the
	// Payee “B” Record.) Please make every effort to compute the
	// taxable amount.
	TaxableAmountNotDeterminedIndicator string `json:"taxable_amount_not_determined_indicator"`

	// Enter “1” (one) for a traditional IRA, SEP, or SIMPLE
	// distribution or Roth conversion. Otherwise, enter a blank. If
	// the IRA/SEP/SIMPLE Indicator is used, enter the amount of
	// the Roth conversion or distribution in Payment Amount Field
	// A of the Payee “B” Record. Do not use the indicator for a
	// distribution from a Roth or for an IRA recharacterization.
	// Note: For Form 1099-R, generally, report the Roth
	// conversion or total amount distributed from a traditional IRA,
	// SEP, or SIMPLE in Payment Amount Field A (traditional
	// IRA/SEP/SIMPLE distribution or Roth conversion), as well as
	// Payment Amount Field 1 (Gross Distribution) of the “B”
	// Record. Refer to Instructions for Forms 1099-R and 5498 for
	// exceptions (Box 2a instructions).
	ISSIndicator string `json:"ira_sep_simple_indicator"`

	// Enter a “1” (one) only if the payment shown for Distribution
	// Amount Code 1 is a total distribution that closed out the
	// account. Otherwise, enter a blank.
	// Note: A total distribution is one or more distributions within
	// one tax year in which the entire balance of the account is
	// distributed. Any distribution that does not meet this definition
	// is not a total distribution.
	TotalDistributionIndicator string `json:"total_distribution_indicator"`

	// Use this field when reporting a total distribution to more than
	// one person, such as when a participant is deceased and a
	// payer distributes to two or more beneficiaries. Therefore, if
	// the percentage is 100, leave this field blank. If the
	// percentage is a fraction, round off to the nearest whole
	// number (for example, 10.4 percent will be10 percent; 10.5
	// percent will be 11 percent). Enter the percentage received by
	// the person whose TIN is included in positions 12-20 of the
	// “B” Record. This field must be right justified, and unused
	// positions must be zero-filled. If not applicable, enter blanks.
	// Filers are not required to enter this information for any IRA
	// distribution or for direct rollovers.
	PercentageTotalDistribution int `json:"percentage_total_distribution"`

	// Enter the first year a designated Roth contribution was made
	// in YYYY format. If the date is unavailable, enter blanks.
	FirstYearDesignatedRothContribution int `json:"firstYear_designated_roth_contribution"`

	// Enter "1" (one) if there is FATCA filing requirement.
	// Otherwise, enter a blank.
	FATCA string `json:"fatca_requirement_indicator"`

	// Enter date of payment in YYYYMMDD format. (for example,
	// January 5, 2019, would be 20190105). Do not enter hyphens
	// or slashes.
	DatePayment time.Time `json:"date_payment"`

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

// Type returns type of “1099-R” record
func (r *Sub1099R) Type() string {
	return config.Sub1099RType
}

// Type returns FS code of “1099-R” record
func (r *Sub1099R) FederalState() int {
	return r.CombinedFSCode
}

// Parse parses the “1099-R” record from fire ascii
func (r *Sub1099R) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099RLayout, record)
}

// Ascii returns fire ascii of “1099-R” record
func (r *Sub1099R) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099RLayout)
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
func (r *Sub1099R) Validate() error {
	return utils.Validate(r, config.Sub1099RLayout, config.Sub1099RType)
}

// Unmarshal parses the JSON-encoded data
func (r *Sub1099R) UnmarshalJSON(data []byte) error {
	type recordJson Sub1099R
	vRecord := recordJson{}
	vRecord.PercentageTotalDistribution = 100
	err := json.Unmarshal(data, &vRecord)
	if err != nil {
		return err
	}
	utils.CopyStruct(&vRecord, r)
	return nil
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1099R) ValidateDistributionCode() error {
	for _, code := range config.DistributionCodes {
		if code == r.DistributionCode {
			return nil
		}
	}
	return utils.NewErrValidValue("distribution code")
}

func (r *Sub1099R) ValidateISSIndicator() error {
	if len(r.ISSIndicator) > 0 &&
		r.ISSIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("ira sep simple indicator")
	}
	return nil
}

func (r *Sub1099R) ValidateTotalDistributionIndicator() error {
	if len(r.TotalDistributionIndicator) > 0 &&
		r.TotalDistributionIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("total distribution indicator")
	}
	return nil
}

func (r *Sub1099R) ValidateTaxableAmountNotDeterminedIndicator() error {
	if len(r.TaxableAmountNotDeterminedIndicator) > 0 &&
		r.TaxableAmountNotDeterminedIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("taxable amount not determined indicator")
	}
	return nil
}

func (r *Sub1099R) ValidateFATCA() error {
	if len(r.FATCA) > 0 &&
		r.FATCA != config.FatcaFilingRequirementIndicator {
		return utils.NewErrValidValue("fatca filing requirement indicator")
	}
	return nil
}

func (r *Sub1099R) ValidateCombinedFSCode() error {
	return utils.ValidateCombinedFSCode(r.CombinedFSCode)
}
