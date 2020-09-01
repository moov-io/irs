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

type Sub1099CAP struct {
	// Enter the date the stock was exchanged for cash, stock in
	// the successor corporation, or other property received in
	// YYYYMMDD format (for example, January 5, 2019, would be
	// 20190105).
	// Do not enter hyphens or slashes.
	DateSaleExchange time.Time `json:"date_sale_exchange"`

	// Enter the number of shares of the corporation’s stock which
	// were exchanged in the transaction.
	// Report whole numbers only. Right justify the information and
	// fill unused positions with zeros.
	NumberSharesExchanged int `json:"number_shares_exchanged"`

	// Enter the class of stock that was exchanged. Left justify the
	// information and fill unused positions with blanks.
	ClassesStockExchanged string `json:"classes_stock_exchanged"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for the filing requirements. If this field is
	// not used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`
}

// Type returns type of “1099-CAP” record
func (r *Sub1099CAP) Type() string {
	return config.Sub1099CapType
}

// Parse parses the “1099-CAP” record from fire ascii
func (r *Sub1099CAP) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099CAPLayout, record)
}

// Ascii returns fire ascii of “1099-CAP” record
func (r *Sub1099CAP) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099CAPLayout)
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
func (r *Sub1099CAP) Validate() error {
	return utils.Validate(r, config.Sub1099CAPLayout, config.Sub1099CapType)
}
