// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"encoding/json"
	"gopkg.in/check.v1"
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
}

func (t *RecordTest) TestKRecordWithError(c *check.C) {
	r := &KRecord{}
	err := r.Parse(t.kRecordAscii[1:])
	c.Assert(err, check.Not(check.IsNil))
}
