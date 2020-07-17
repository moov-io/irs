// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

// General record interface
type Record interface {
	Type() string
	SequenceNumber() int
	SetSequenceNumber(int)
	Parse([]byte) error
	Ascii() []byte
	Validate() error
}
