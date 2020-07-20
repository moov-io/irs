// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package config

const (
	// TRecordType indicates type of transmitter “T” record
	TRecordType = "T"
	// ARecordType indicates type of payer “A” record
	ARecordType = "A"
	// BRecordType indicates type of payee “B” record
	BRecordType = "B"
	// CRecordType indicates type of end payer “C” record
	CRecordType = "C"
	// KRecordType indicates type of totals “K” record
	KRecordType = "K"
	// FRecordType indicates name of transmission “F” record
	FRecordType = "F"
	// Sub1097BtcType indicates type of payee “B” record for form 1097-BTC
	Sub1097BtcType = "1097-BTC"
	// Sub1099IntType indicates type of payee “B” record for form 1099-INT
	Sub1099IntType = "1099-INT"
	// Sub1099MiscType indicates type of payee “B” record for form 1099-MISC
	Sub1099MiscType = "1099-MISC"
	// Sub1099OidType indicates type of payee “B” record for form 1099-OID
	Sub1099OidType = "1099-OID"
	// Sub1099PatrType indicates type of payee “B” record for form 1099-PATR
	Sub1099PatrType = "1099-PATR"
)

const (
	// RecordLength indicates length of general record
	RecordLength = 750
	// SubRecordLength indicates length of sub record for payee “B” record
	SubRecordLength = 207
)

const (
	// BlankString indicates the empty string
	BlankString = " "
	// ZeroString indicates the zero string
	ZeroString = "0"
)
