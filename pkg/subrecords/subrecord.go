// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package subrecords

// General subrecord interface
type SubRecord interface {
	Type() string
	Parse([]byte) error
	Ascii() []byte
	Validate() error
}
