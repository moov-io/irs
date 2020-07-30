// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"encoding/json"
	"gopkg.in/check.v1"

	"github.com/moov-io/irs/pkg/config"
)

func (t *RecordTest) TestKRecord(c *check.C) {
	r := NewKRecord()
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.kRecordJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.kRecordAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.kRecordAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.kRecordAscii))
	r.SetSequenceNumber(1)
	c.Assert(r.SequenceNumber(), check.Equals, 1)
	c.Assert(r.Validate(), check.IsNil)
	r.SetSequenceNumber(-1)
	c.Assert(r.Validate(), check.NotNil)
	c.Assert(r.Type(), check.Equals, config.KRecordType)
	kRecord := r.(*KRecord)
	codes := kRecord.PaymentCodes()
	c.Assert(len(codes), check.Not(check.Equals), 0)
	_, err = kRecord.ControlTotal("1")
	c.Assert(err, check.IsNil)
}

func (t *RecordTest) TestKRecordWithError(c *check.C) {
	r := &KRecord{}
	err := r.Parse(t.kRecordAscii[1:])
	c.Assert(err, check.Not(check.IsNil))
}
