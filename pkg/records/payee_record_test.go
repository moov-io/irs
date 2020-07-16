// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"encoding/json"
	"gopkg.in/check.v1"
)

func (t *RecordTest) TestBRecord(c *check.C) {
	r := &BRecord{}
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.bRecordJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecordAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecordAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecordAscii))
}

func (t *RecordTest) TestBRecordWithError(c *check.C) {
	r := &BRecord{}
	err := r.Parse(t.bRecordAscii[1:])
	c.Assert(err, check.Not(check.IsNil))
}
