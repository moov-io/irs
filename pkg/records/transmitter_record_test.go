// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"encoding/json"
	"gopkg.in/check.v1"

	"github.com/moov-io/irs/pkg/config"
)

func (t *RecordTest) TestTRecord(c *check.C) {
	r := NewTRecord()
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.tRecordJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.tRecordAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.tRecordAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.tRecordAscii))
	c.Assert(r.Type(), check.Equals, config.TRecordType)
	r.SetSequenceNumber(1)
	c.Assert(r.SequenceNumber(), check.Equals, 1)
	c.Assert(r.Validate(), check.IsNil)
	r.SetSequenceNumber(-1)
	c.Assert(r.Validate(), check.NotNil)
}

func (t *RecordTest) TestTRecordWithError(c *check.C) {
	r := &TRecord{}
	err := r.Parse(t.tRecordAscii[1:])
	c.Assert(err, check.Not(check.IsNil))
	err = r.Parse(t.tRecordAscii)
	c.Assert(err, check.IsNil)
	r.VendorForeignEntityIndicator = "ERR"
	c.Assert(r.ValidateVendorForeignEntityIndicator(), check.Not(check.IsNil))
	r.VendorState = "ERR"
	c.Assert(r.ValidateVendorState(), check.Not(check.IsNil))
	r.VendorIndicator = "ERR"
	c.Assert(r.ValidateVendorIndicator(), check.Not(check.IsNil))
	r.CompanyState = "ERR"
	c.Assert(r.ValidateCompanyState(), check.Not(check.IsNil))
	r.ForeignEntityIndicator = "ERR"
	c.Assert(r.ValidateForeignEntityIndicator(), check.Not(check.IsNil))
	r.TestFileIndicator = "ERR"
	c.Assert(r.ValidateTestFileIndicator(), check.Not(check.IsNil))
	r.PriorYearDataIndicator = "ERR"
	c.Assert(r.ValidatePriorYearDataIndicator(), check.Not(check.IsNil))
}

func (t *RecordTest) TestTRecord_BlankFields(c *check.C) {
	r := NewTRecord()
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.tRecordJson, r)
	c.Assert(err, check.IsNil)
	transmitter := r.(*TRecord)
	transmitter.PriorYearDataIndicator = ""
	transmitter.TestFileIndicator = ""
	transmitter.ForeignEntityIndicator = ""
	transmitter.VendorForeignEntityIndicator = ""
	c.Assert(r.Validate(), check.IsNil)
}
