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
	// Sub1097BTCType indicates type of payee “B” record for form 1097-BTC
	Sub1097BTCType = "1097-BTC"
	// Sub1099INTType indicates type of payee “B” record for form 1099-INT
	Sub1099INTType = "1099-INT"
	// Sub1099MISCType indicates type of payee “B” record for form 1099-MISC
	Sub1099MISCType = "1099-MISC"
	// Sub1099OIDType indicates type of payee “B” record for form 1099-OID
	Sub1099OIDType = "1099-OID"
	// Sub1099OIDType indicates type of payee “B” record for form 1099-PATR
	Sub1099PATRType = "1099-PATR"
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
