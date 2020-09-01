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

type Sub1098T struct {
	// Required.
	// Enter 1 to certify compliance with applicable TIN solicitation
	// requirements regarding individual student when:
	// • Educational institution received a TIN from the
	//   individual in response to specific solicitation in the
	//   current year, a previous year, or the institution
	//   obtained the TIN from the student’s application for
	//   financial aid or other form (whether in the year for
	//   which the form is filed or a prior year) and, in either
	//   instance, has no reason to believe the TIN on file in
	//   the institution’s records is incorrect.
	// • Educational institution files Form 1098-T with this
	// field blank because it has no record of the student’s
	// TIN, but only if the institution made the required
	// written TIN solicitation by December 31 of the
	// calendar year for which the Form 1098-T is being
	// filed.
	// Otherwise, leave blank.
	IdentificationNumber string `json:"identification_number"`

	// Required. Enter “1” (one) if the student was at least a halftime student during any academic period that began in 2019.
	// Otherwise, enter a blank.
	HalfTimeStudentIndicator string `json:"halftime_student_indicator"`

	// Required. Enter “1” (one) if the student is enrolled
	// exclusively in a graduate level program. Otherwise, enter a
	// blank.
	GraduateStudentIndicator string `json:"graduate_student_indicator"`

	// Enter “1” (one) if the amount in Payment Amount Field 1 or
	// Payment Amount Field 2 includes amounts for an academic
	// period beginning January through March 2020. Otherwise,
	// enter a blank.
	AcademicPeriodIndicator string `json:"academic_period_indicator"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or
	// local revenue departments for the filing requirements. If
	// this field is not use, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “1098-T” record
func (r *Sub1098T) Type() string {
	return config.Sub1098TType
}

// Parse parses the “1098-T” record from fire ascii
func (r *Sub1098T) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1098TLayout, record)
}

// Ascii returns fire ascii of “1098-T” record
func (r *Sub1098T) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1098TLayout)
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
func (r *Sub1098T) Validate() error {
	return utils.Validate(r, config.Sub1098TLayout, config.Sub1098TType)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1098T) ValidateIdentificationNumber() error {
	if len(r.IdentificationNumber) > 0 &&
		r.IdentificationNumber != config.GeneralOneIndicator {
		return utils.NewErrValidValue("identification number")
	}
	return nil
}

func (r *Sub1098T) ValidateHalfTimeStudentIndicator() error {
	if len(r.HalfTimeStudentIndicator) > 0 &&
		r.HalfTimeStudentIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("half-time student indicator")
	}
	return nil
}

func (r *Sub1098T) ValidateGraduateStudentIndicator() error {
	if len(r.GraduateStudentIndicator) > 0 &&
		r.GraduateStudentIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("graduate student indicator")
	}
	return nil
}

func (r *Sub1098T) ValidateAcademicPeriodIndicator() error {
	if len(r.AcademicPeriodIndicator) > 0 &&
		r.AcademicPeriodIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("academic period indicator")
	}
	return nil
}
