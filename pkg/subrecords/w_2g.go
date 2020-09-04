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

type SubW2G struct {
	// Required. Enter the applicable type of wager code from the
	// table below.
	// 1: Horse race track
	// 2: Dog race track
	// 3: Jai-alai
	// 4: State-conducted lottery
	// 5: Keno
	// 6: Bingo
	// 7: Slot machines
	// 8: Poker winnings
	// 9: Any other type of gambling winnings
	TypeWagerCode string `json:"type_wager_code"`

	// Required. Enter the date of the winning transaction in
	// YYYYMMDD format (for example, January 5, 2019, would be
	// 20190105). This is not the date the money was paid, if paid
	// after the date of the race (or game).
	// Do not enter hyphens or slashes.
	DateWon time.Time `json:"date_won"`

	// Required. For state-conducted lotteries, enter the ticket or
	// other identifying number.
	// For keno, bingo, and slot machines, enter the ticket or card
	// number (and color, if applicable), machine serial number, or
	// any other information that will help identify the winning
	// transaction.
	// For all others, enter blanks.
	Transaction string `json:"transaction"`

	// If applicable, enter the race (or game) relating to the winning
	// ticket. Otherwise, enter blanks.
	Race string `json:"race"`

	// If applicable, enter the initials or number of the cashier
	// making the winning payment. Otherwise, enter blanks.
	Cashier string `json:"cashier"`

	// If applicable, enter the window number or location of the
	// person paying the winning payment. Otherwise, enter
	// blanks.
	Window string `json:"window"`

	// For other than state lotteries, enter the first identification
	// number of the person receiving the winning payment.
	// Otherwise, enter blanks.
	FirstID string `json:"first_id"`

	// For other than state lotteries, enter the second identification
	// number of the person receiving the winnings. Otherwise,
	// enter blanks.
	SecondID string `json:"second_id"`

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

// Type returns type of “W-2G” record
func (r *SubW2G) Type() string {
	return config.SubW2GType
}

// Type returns FS code of “W-2G” record
func (r *SubW2G) FederalState() int {
	return 0
}

// Parse parses the “W-2G” record from fire ascii
func (r *SubW2G) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.SubW2GLayout, record)
}

// Ascii returns fire ascii of “W-2G” record
func (r *SubW2G) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.SubW2GLayout)
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
func (r *SubW2G) Validate() error {
	return utils.Validate(r, config.SubW2GLayout, config.SubW2GType)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *SubW2G) ValidateTypeWagerCode() error {
	if r.TypeWagerCode[0] >= '1' && r.TypeWagerCode[0] <= '2' {
		return nil
	}
	return utils.NewErrValidValue("type wager code")
}
