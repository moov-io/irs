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
	err := r.SetTypeOfReturn(config.Sub1099MiscType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099MiscJson, r)
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
	codes := r.PaymentCodes()
	c.Assert(len(codes), check.Not(check.Equals), 0)
	_, err = r.PaymentAmount("1")
	c.Assert(err, check.IsNil)
}

func (t *RecordTest) TestBRecordWithError(c *check.C) {
	r := &BRecord{}
	err := r.Parse(t.bRecord1099MiscAscii[1:])
	c.Assert(err, check.Not(check.IsNil))
	err = r.Parse(t.bRecord1099MiscAscii)
	c.Assert(err, check.IsNil)
	err = r.Validate()
	c.Assert(err.Error(), check.Equals, "should exist extension block")
	_, err = r.PaymentAmount("K")
	c.Assert(err.Error(), check.Equals, "is an invalid field")
	c.Assert("", check.Equals, r.TypeOfReturn())
	_, err = json.Marshal(r)
	c.Assert(err, check.IsNil)
	r.CorrectedReturnIndicator = "K"
	err = r.ValidateCorrectedReturnIndicator()
	c.Assert(err, check.Not(check.IsNil))
	r.TypeOfTIN = "3"
	err = r.ValidateTypeOfTIN()
	c.Assert(err, check.Not(check.IsNil))
	r.ForeignCountryIndicator = "3"
	err = r.ValidateForeignCountryIndicator()
	c.Assert(err, check.Not(check.IsNil))
	r.PayeeState = "3"
	err = r.ValidatePayeeState()
	c.Assert(err, check.Not(check.IsNil))
	err = json.Unmarshal([]byte("error string"), r)
	c.Assert(err, check.Not(check.IsNil))
}

func (t *RecordTest) TestBRecordWith1099Int(c *check.C) {
	r, err := NewBRecord(config.Sub1099IntType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099IntJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.(*BRecord).extRecord.Type(), check.Equals, config.Sub1099IntType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099IntAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099IntAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099IntAscii))
}

func (t *RecordTest) TestBRecordWith1099Oid(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099OidType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099OidJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099OidType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099OidAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099OidAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099OidAscii))
}

func (t *RecordTest) TestBRecordWith1099Patr(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099PatrType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099PatrJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099PatrType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099PatrAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099PatrAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099PatrAscii))
}

func (t *RecordTest) TestBRecordWith1097Btc(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1097BtcType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1097BtcJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1097BtcType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1097BtcAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1097BtcAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1097BtcAscii))
}

func (t *RecordTest) TestBRecordWith1098(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1098Type)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1098Json, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1098Type)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1098Ascii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1098Ascii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1098Ascii))
}

func (t *RecordTest) TestBRecordWith1098C(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1098CType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1098CJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1098CType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1098CAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1098CAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1098CAscii))
}

func (t *RecordTest) TestBRecordWith1098E(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1098EType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1098EJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1098EType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1098EAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1098EAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1098EAscii))
}

func (t *RecordTest) TestBRecordWith1098F(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1098FType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1098FJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1098FType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1098FAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1098FAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1098FAscii))
}

func (t *RecordTest) TestBRecordWith1098Q(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1098QType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1098QJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1098QType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1098QAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1098QAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1098QAscii))
}

func (t *RecordTest) TestBRecordWith1098T(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1098TType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1098TJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1098TType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1098TAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1098TAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1098TAscii))
}
