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

type Sub1098Q struct {
	// Enter the annuity start date in YYYYMMDD format. If the
	// payments have not started, show the annuity amount payable on
	// start date in YYYYMMDD format.
	AnnuityStartDate time.Time `json:"annuity_start_date"`

	// Enter “1” (one) if payments have not yet started and the start
	// date may be accelerated. Otherwise, enter a blank.
	AcceleratedIndicator string `json:"accelerated_indicator"`

	// Enter a two-digit number 01-31. Otherwise, enter blanks.
	January int `json:"january"`

	// Enter a two-digit number 01-31. Otherwise, enter blanks.
	February int `json:"february"`

	// Enter a two-digit number 01-31. Otherwise, enter blanks.
	March int `json:"march"`

	// Enter a two-digit number 01-31. Otherwise, enter blanks.
	April int `json:"april"`

	// Enter a two-digit number 01-31. Otherwise, enter blanks.
	May int `json:"may"`

	// Enter a two-digit number 01-31. Otherwise, enter blanks.
	June int `json:"june"`

	// Enter a two-digit number 01-31. Otherwise, enter blanks.
	July int `json:"july"`

	// Enter a two-digit number 01-31. Otherwise, enter blanks.
	August int `json:"august"`

	// Enter a two-digit number 01-31. Otherwise, enter blanks.
	September int `json:"september"`

	// Enter a two-digit number 01-31. Otherwise, enter blanks.
	October int `json:"october"`

	// Enter a two-digit number 01-31. Otherwise, enter blanks.
	November int `json:"november"`

	// Enter a two-digit number 01-31. Otherwise, enter blanks.
	December int `json:"december"`

	// If the contract was purchased under a plan, enter the name of
	// the plan. Otherwise, enter blanks.
	NamePlan string `json:"name_plan"`

	// If the contract was purchased under a plan, enter the plan
	// number. Otherwise, enter blanks.
	PlanNumber string `json:"plan_number"`

	// If the contract was purchased under a plan, enter the nine-digit
	// employer identification number of the plan sponsor. Otherwise,
	// enter blanks.
	EmployerIdentificationNumber string `json:"employer_identification_number"`
}

// Type returns type of “1098-Q” record
func (r *Sub1098Q) Type() string {
	return config.Sub1098QType
}

// Parse parses the “1098-Q” record from fire ascii
func (r *Sub1098Q) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1098QLayout, record)
}

// Ascii returns fire ascii of “1098-Q” record
func (r *Sub1098Q) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1098QLayout)
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
func (r *Sub1098Q) Validate() error {
	return utils.Validate(r, config.Sub1098QLayout)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1098Q) ValidateAcceleratedIndicator() error {
	if len(r.AcceleratedIndicator) > 0 &&
		r.AcceleratedIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("accelerated indicator")
	}
	return nil
}

func (r *Sub1098Q) ValidateJanuary() error {
	if r.January >= 0 && r.January <= 31 {
		return nil
	}
	return utils.NewErrValidValue("january")
}

func (r *Sub1098Q) ValidateFebruary() error {
	if r.February >= 0 && r.February <= 28 {
		return nil
	}
	return utils.NewErrValidValue("february")
}

func (r *Sub1098Q) ValidateMarch() error {
	if r.March >= 0 && r.March <= 31 {
		return nil
	}
	return utils.NewErrValidValue("march")
}

func (r *Sub1098Q) ValidateApril() error {
	if r.April >= 0 && r.April <= 30 {
		return nil
	}
	return utils.NewErrValidValue("april")
}

func (r *Sub1098Q) ValidateMay() error {
	if r.May >= 0 && r.May <= 31 {
		return nil
	}
	return utils.NewErrValidValue("may")
}

func (r *Sub1098Q) ValidateJune() error {
	if r.June >= 0 && r.June <= 30 {
		return nil
	}
	return utils.NewErrValidValue("june")
}

func (r *Sub1098Q) ValidateJuly() error {
	if r.July >= 0 && r.July <= 31 {
		return nil
	}
	return utils.NewErrValidValue("july")
}

func (r *Sub1098Q) ValidateAugust() error {
	if r.August >= 0 && r.August <= 31 {
		return nil
	}
	return utils.NewErrValidValue("august")
}

func (r *Sub1098Q) ValidateSeptember() error {
	if r.September >= 0 && r.September <= 30 {
		return nil
	}
	return utils.NewErrValidValue("september")
}

func (r *Sub1098Q) ValidateOctober() error {
	if r.October >= 0 && r.October <= 31 {
		return nil
	}
	return utils.NewErrValidValue("october")
}

func (r *Sub1098Q) ValidateNovember() error {
	if r.November >= 0 && r.November <= 31 {
		return nil
	}
	return utils.NewErrValidValue("november")
}

func (r *Sub1098Q) ValidateDecember() error {
	if r.December >= 0 && r.December <= 31 {
		return nil
	}
	return utils.NewErrValidValue("december")
}

func (r *Sub1098Q) ValidateEmployerIdentificationNumber() error {
	if len(r.EmployerIdentificationNumber) >= 0 {
		return utils.IsNumeric(r.EmployerIdentificationNumber)
	}
	return utils.NewErrValidValue("employer identification number")
}
