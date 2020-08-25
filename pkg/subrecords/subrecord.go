// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package subrecords

import (
	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/utils"
)

// General subrecord interface
type SubRecord interface {
	Type() string
	Parse([]byte) error
	Ascii() []byte
	Validate() error
}

// NewSubRecord returns a new sub record with type of return
func NewSubRecord(recordType string) (SubRecord, error) {
	var newRecord SubRecord
	switch recordType {
	case config.Sub1097BtcType:
		newRecord = &Sub1097BTC{}
	case config.Sub1098Type:
		newRecord = &Sub1098{}
	case config.Sub1098CType:
		newRecord = &Sub1098C{}
	case config.Sub1098EType:
		newRecord = &Sub1098E{}
	case config.Sub1098FType:
		newRecord = &Sub1098F{}
	case config.Sub1098QType:
		newRecord = &Sub1098Q{}
	case config.Sub1098TType:
		newRecord = &Sub1098T{}
	case config.Sub1099AType:
		newRecord = &Sub1099A{}
	case config.Sub1099BType:
		newRecord = &Sub1099B{}
	case config.Sub1099CType:
		newRecord = &Sub1099C{}
	case config.Sub1099CapType:
		newRecord = &Sub1099CAP{}
	case config.Sub1099DivType:
		newRecord = &Sub1099DIV{}
	case config.Sub1099GType:
		newRecord = &Sub1099G{}
	case config.Sub1099HType:
		newRecord = &Sub1099H{}
	case config.Sub1099IntType:
		newRecord = &Sub1099INT{}
	case config.Sub1099KType:
		newRecord = &Sub1099K{}
	case config.Sub1099LsType:
		newRecord = &Sub1099LS{}
	case config.Sub1099LtcType:
		newRecord = &Sub1099LTC{}
	case config.Sub1099MiscType:
		newRecord = &Sub1099MISC{}
	case config.Sub1099OidType:
		newRecord = &Sub1099OID{}
	default:
		return newSubRecord(recordType)
	}
	return newRecord, nil
}

func newSubRecord(recordType string) (SubRecord, error) {
	var newRecord SubRecord
	switch recordType {
	case config.Sub1099PatrType:
		newRecord = &Sub1099PATR{}
	case config.Sub1099QType:
		newRecord = &Sub1099Q{}
	case config.Sub1099RType:
		newRecord = &Sub1099R{}
	case config.Sub1099SType:
		newRecord = &Sub1099S{}
	case config.Sub1099SaType:
		newRecord = &Sub1099SA{}
	case config.Sub1099SbType:
		newRecord = &Sub1099SB{}
	default:
		return nil, utils.ErrUnsupportedBlock
	}
	return newRecord, nil
}
