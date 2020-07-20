// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package subrecords

import "github.com/moov-io/irs/pkg/config"

// General subrecord interface
type SubRecord interface {
	Type() string
	Parse([]byte) error
	Ascii() []byte
	Validate() error
}

// NewSubRecord returns a new sub record with type of return
func NewSubRecord(recordType string) SubRecord {
	var newRecord SubRecord = nil
	switch recordType {
	case config.Sub1097BtcType:
		newRecord = &Sub1097BTC{}
	case config.Sub1099IntType:
		newRecord = &Sub1099INT{}
	case config.Sub1099MiscType:
		newRecord = &Sub1099MISC{}
	case config.Sub1099OidType:
		newRecord = &Sub1099OID{}
	case config.Sub1099PatrType:
		newRecord = &Sub1099PATR{}
	}
	return newRecord
}
