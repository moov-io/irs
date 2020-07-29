// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"bytes"
	"encoding/json"
	"gopkg.in/check.v1"

	"github.com/moov-io/irs/pkg/config"
)

func (t *RecordTest) TestBRecordWith1099MISC(c *check.C) {
	r := &BRecord{}
	r.SetTypeOfReturn(config.Sub1099MiscType)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.bRecord1099MiscJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099MiscAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099MiscAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099MiscAscii))
	buf, err := json.Marshal(r)
	c.Assert(err, check.IsNil)
	var prettyJSON1 bytes.Buffer
	json.Indent(&prettyJSON1, buf, "", "  ")
	var prettyJSON2 bytes.Buffer
	json.Indent(&prettyJSON2, t.bRecord1099MiscJson, "", "  ")
	c.Assert(prettyJSON1.String(), check.Equals, prettyJSON2.String())
	r.SetSequenceNumber(1)
	c.Assert(r.SequenceNumber(), check.Equals, 1)
	c.Assert(r.Validate(), check.IsNil)
	r.SetSequenceNumber(-1)
	c.Assert(r.Validate(), check.NotNil)
	c.Assert(r.Type(), check.Equals, config.BRecordType)
}

func (t *RecordTest) TestBRecordWithError(c *check.C) {
	r := &BRecord{}
	err := r.Parse(t.bRecord1099MiscAscii[1:])
	c.Assert(err, check.Not(check.IsNil))
}

func (t *RecordTest) TestBRecordWith1099Int(c *check.C) {
	r := NewBRecord(config.Sub1099IntType)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.bRecord1099IntJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099IntAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099IntAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099IntAscii))
}

func (t *RecordTest) TestBRecordWith1099Oid(c *check.C) {
	r := &BRecord{}
	r.SetTypeOfReturn(config.Sub1099OidType)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.bRecord1099OidJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099OidAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099OidAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099OidAscii))
}

func (t *RecordTest) TestBRecordWith1099Patr(c *check.C) {
	r := &BRecord{}
	r.SetTypeOfReturn(config.Sub1099PatrType)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.bRecord1099PatrJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099PatrAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099PatrAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099PatrAscii))
}

func (t *RecordTest) TestBRecordWith1097Btc(c *check.C) {
	r := &BRecord{}
	r.SetTypeOfReturn(config.Sub1097BtcType)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.bRecord1097BtcJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1097BtcAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1097BtcAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1097BtcAscii))
}
