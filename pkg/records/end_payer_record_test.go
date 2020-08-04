// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"encoding/json"
	"gopkg.in/check.v1"

	"github.com/moov-io/irs/pkg/config"
)

func (t *RecordTest) TestCRecord(c *check.C) {
	r := NewCRecord()
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.cRecordJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.cRecordAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.cRecordAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.cRecordAscii))
	r.SetSequenceNumber(1)
	c.Assert(r.SequenceNumber(), check.Equals, 1)
	c.Assert(r.Validate(), check.IsNil)
	r.SetSequenceNumber(-1)
	c.Assert(r.Validate(), check.NotNil)
	c.Assert(r.Type(), check.Equals, config.CRecordType)
	cRecord := r.(*CRecord)
	codes := cRecord.TotalCodes()
	c.Assert(len(codes), check.Not(check.Equals), 0)
	_, err = cRecord.ControlTotal("1")
	c.Assert(err, check.IsNil)
}

func (t *RecordTest) TestCRecordWithError(c *check.C) {
	r := &CRecord{}
	err := r.Parse(t.cRecordAscii[1:])
	c.Assert(err, check.Not(check.IsNil))
	err = r.Parse(t.cRecordAscii)
	c.Assert(err, check.IsNil)
	_, err = r.ControlTotal("K")
	c.Assert(err, check.Not(check.IsNil))
}
