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

type Sub1099SA struct {
	// Required. Enter the applicable code from the table below to
	// indicate the type of payment.
	// 1: Normal distribution
	// 2: Excess contribution
	// 3: Disability
	// 4: Death distribution other than code 6
	// 5: Prohibited transaction
	// 6: Death distribution after the year of death to a nonspouse beneficiary.
	DistributionCode string `json:"distribution_code"`

	// Enter “1” (one) if distributions are from a Medicare Advantage
	// MSA. Otherwise, enter a blank.
	MedicareAdvantageMSAIndicator string `json:"medicare_advantage_msa_indicator"`

	// Enter “1” (one) if distributions are from a HAS. Otherwise,
	// enter a blank.
	HSAIndicator string `json:"hsa_indicator"`

	// Enter “1” (one) if distributions are from an Archer MSA.
	// Otherwise, enter a blank.
	ArcherMSAIndicator string `json:"archer_mas_indicator"`

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
}

// Type returns type of “1099-SA” record
func (r *Sub1099SA) Type() string {
	return config.Sub1099SaType
}

// Parse parses the “1099-SA” record from fire ascii
func (r *Sub1099SA) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099SALayout, record)
}

// Ascii returns fire ascii of “1099-SA” record
func (r *Sub1099SA) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099SALayout)
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
func (r *Sub1099SA) Validate() error {
	return utils.Validate(r, config.Sub1099SALayout)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1099SA) ValidateDistributionCode() error {
	switch r.DistributionCode {
	case "1", "2", "3", "4", "5", "6":
		return nil
	}
	return utils.NewErrValidValue("distribution code")
}

func (r *Sub1099SA) ValidateMedicareAdvantageMSAIndicator() error {
	if len(r.MedicareAdvantageMSAIndicator) > 0 &&
		r.MedicareAdvantageMSAIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("medicare advantage msa indicator")
	}
	return nil
}

func (r *Sub1099SA) ValidateHSAIndicator() error {
	if len(r.HSAIndicator) > 0 &&
		r.HSAIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("has indicator")
	}
	return nil
}

func (r *Sub1099SA) ValidateArcherMSAIndicator() error {
	if len(r.ArcherMSAIndicator) > 0 &&
		r.ArcherMSAIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("archer msa indicator")
	}
	return nil
}
