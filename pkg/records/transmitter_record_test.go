// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"encoding/json"
	"gopkg.in/check.v1"
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
}

func (t *RecordTest) TestTRecordWithError(c *check.C) {
	r := &TRecord{}
	err := r.Parse(t.tRecordAscii[1:])
	c.Assert(err, check.Not(check.IsNil))
}
