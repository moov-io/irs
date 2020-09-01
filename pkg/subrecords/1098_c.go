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

type Sub1098C struct {
	// Enter “1” (one) if the amount reported in Payment Amount
	// Field 4 is an arm’s length transaction to an unrelated party.
	// Otherwise, enter a blank.
	TransactionIndicator string `json:"transaction_indicator"`

	// Enter “1” (one) if the vehicle will not be transferred for
	// money, other property, or services before completion of
	// material improvements or significant intervening use.
	// Otherwise, enter a blank.
	TransferAfterImprovementsIndicator string `json:"transfer_after_improvements_indicator"`

	// Enter “1” (one) if the vehicle is transferred to a needy
	// individual for significantly below fair market value. Otherwise,
	// enter a blank.
	TransferMarketValueIndicator string `json:"transfer_market_value_indicator"`

	// Enter the year of the vehicle in YYYY format.
	Year int `json:"year"`

	// Enter the Make of the vehicle. Left justify the information and
	// fill unused positions with blanks.
	Make string `json:"make"`

	// Enter the Model of the vehicle. Left justify the information and
	// fill unused positions with blanks.
	Model string `json:"model"`

	// Enter the vehicle or other identification number of the
	// donated vehicle. Left justify the information and fill unused
	// positions with blanks.
	VehicleIdentificationNumber string `json:"vehicle_identification_number"`

	// Enter a description of material improvements or significant
	// intervening use and duration of use. Left justify the
	// information and fill unused positions with blanks.
	VehicleDescription string `json:"vehicle_description"`

	// Enter the date the contribution was made to an organization,
	// in YYYYMMDD format (for example, January 5, 2019, would
	// be 20190105).
	DateContribution time.Time `json:"date_contribution"`

	// Enter the appropriate indicator from the following table to
	// report if the donee of the vehicle provides goods or services
	// in exchange for the vehicle.
	// 1: Donee provided goods or services
	// 2: Donee did not provide goods or services
	DoneeIndicator string `json:"donee_indicator"`

	// Enter “1” (one) if only intangible religious benefits were
	// provided in exchange for the vehicle. Otherwise, enter a
	// blank.
	IntangibleReligiousBenefitsIndicator string `json:"intangible_religious_benefits_indicator"`

	// Enter “1” (one) if under the law the donor cannot claim a
	// deduction of more than $500 for the vehicle. Otherwise,
	// enter a blank.
	DeductionLessIndicator string `json:"deduction_less_indicator"`

	// You may enter odometer mileage here. Enter as 7 numeric
	// characters. The remaining positions of this field may be used
	// to record information for state and local government
	// reporting or for the filer's own purposes. Payers should
	// contact the state or local revenue departments for the filing
	// requirements. If this field is not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`

	// Enter the date of sale, in YYYYMMDD format (for example,
	// January 5, 2019, would be 20190105). Do not enter hyphens
	// or slashes.
	DateSale time.Time `json:"date_sale"`

	// Enter a description of any goods and services received for
	// the vehicle. Otherwise, enter blanks.
	// Left justify information and fill unused positions with blanks.
	GoodsServices string `json:"goods_services"`
}

// Type returns type of “1098-C” record
func (r *Sub1098C) Type() string {
	return config.Sub1098CType
}

// Parse parses the “1098-C” record from fire ascii
func (r *Sub1098C) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1098CLayout, record)
}

// Ascii returns fire ascii of “1098-C” record
func (r *Sub1098C) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1098CLayout)
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
func (r *Sub1098C) Validate() error {
	return utils.Validate(r, config.Sub1098CLayout, config.Sub1098CType)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1098C) ValidateTransactionIndicator() error {
	if len(r.TransactionIndicator) > 0 && r.TransactionIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("transaction indicator")
	}
	return nil
}

func (r *Sub1098C) ValidateTransferAfterImprovementsIndicator() error {
	if len(r.TransferAfterImprovementsIndicator) > 0 && r.TransferAfterImprovementsIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("transfer after improvements indicator")
	}
	return nil
}

func (r *Sub1098C) ValidateTransferMarketValueIndicator() error {
	if len(r.TransferMarketValueIndicator) > 0 && r.TransferMarketValueIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("transfer market value indicator")
	}
	return nil
}

func (r *Sub1098C) ValidateDoneeIndicator() error {
	if len(r.DoneeIndicator) > 0 && (r.DoneeIndicator != config.GeneralOneIndicator && r.DoneeIndicator != config.GeneralTwoIndicator) {
		return utils.NewErrValidValue("donee indicator")
	}
	return nil
}

func (r *Sub1098C) ValidateIntangibleReligiousBenefitsIndicator() error {
	if len(r.IntangibleReligiousBenefitsIndicator) > 0 && r.IntangibleReligiousBenefitsIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("intangible religious benefits indicator")
	}
	return nil
}

func (r *Sub1098C) ValidateDeductionLessIndicator() error {
	if len(r.DeductionLessIndicator) > 0 && r.DeductionLessIndicator != config.GeneralOneIndicator {
		return utils.NewErrValidValue("deduction less indicator")
	}
	return nil
}
