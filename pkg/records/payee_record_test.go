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
	_, err = r.Fatca()
	c.Assert(err, check.IsNil)
	_, err = r.SecondTIN()
	c.Assert(err, check.IsNil)
	_, err = r.DirectSales()
	c.Assert(err, check.IsNil)
	_, _, err = r.IncomeTax()
	c.Assert(err, check.IsNil)
}

func (t *RecordTest) TestBRecordWith1099NEC(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099NecType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099NecJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099NecAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099NecAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099NecAscii))
	buf, err := json.Marshal(r)
	c.Assert(err, check.IsNil)
	var prettyJSON1 bytes.Buffer
	json.Indent(&prettyJSON1, buf, "", "  ")
	var prettyJSON2 bytes.Buffer
	json.Indent(&prettyJSON2, t.bRecord1099NecJson, "", "  ")
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
	_, err = r.Fatca()
	c.Assert(err, check.IsNil)
	_, err = r.SecondTIN()
	c.Assert(err, check.IsNil)
	_, err = r.DirectSales()
	c.Assert(err, check.NotNil)
	_, _, err = r.IncomeTax()
	c.Assert(err, check.NotNil)
}

func (t *RecordTest) TestBRecordWithError(c *check.C) {
	r := &BRecord{}
	err := r.Parse(t.bRecord1099MiscAscii[1:])
	c.Assert(err, check.NotNil)
	_, err = r.Fatca()
	c.Assert(err, check.NotNil)
	_, err = r.SecondTIN()
	c.Assert(err, check.NotNil)
	_, err = r.DirectSales()
	c.Assert(err, check.NotNil)
	_, _, err = r.IncomeTax()
	c.Assert(err, check.NotNil)
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
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099IntType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099IntJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099IntType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099IntAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099IntAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099IntAscii))
	c.Assert(err, check.IsNil)
	_, err = r.Fatca()
	c.Assert(err, check.NotNil)
	_, err = r.SecondTIN()
	c.Assert(err, check.NotNil)
	_, err = r.DirectSales()
	c.Assert(err, check.NotNil)
	_, _, err = r.IncomeTax()
	c.Assert(err, check.NotNil)
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

func (t *RecordTest) TestBRecordWith1099A(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099AType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099AJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099AType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099AAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099AAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099AAscii))
}

func (t *RecordTest) TestBRecordWith1099B(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099BType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099BJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099BType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099BAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099BAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099BAscii))
}

func (t *RecordTest) TestBRecordWith1099C(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099CType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099CJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099CType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099CAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099CAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099CAscii))
}

func (t *RecordTest) TestBRecordWith1099Cap(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099CapType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099CapJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099CapType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099CapAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099CapAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099CapAscii))
}

func (t *RecordTest) TestBRecordWith1099Div(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099DivType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099DivJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099DivType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099DivAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099DivAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099DivAscii))
}

func (t *RecordTest) TestBRecordWith1099G(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099GType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099GJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099GType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099GAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099GAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099GAscii))
}

func (t *RecordTest) TestBRecordWith1099H(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099HType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099HJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099HType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099HAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099HAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099HAscii))
}

func (t *RecordTest) TestBRecordWith1099K(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099KType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099KJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099KType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099KAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099KAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099KAscii))
}

func (t *RecordTest) TestBRecordWith1099Ls(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099LsType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099LsJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099LsType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099LsAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099LsAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099LsAscii))
}

func (t *RecordTest) TestBRecordWith1099Ltc(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099LtcType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099LtcJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099LtcType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099LtcAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099LtcAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099LtcAscii))
}

func (t *RecordTest) TestBRecordWith1099Q(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099QType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099QJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099QType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099QAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099QAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099QAscii))
}

func (t *RecordTest) TestBRecordWith1099R(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099RType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099RJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099RType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099RAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099RAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099RAscii))
}

func (t *RecordTest) TestBRecordWith1099S(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099SType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099SJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099SType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099SAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099SAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099SAscii))
}

func (t *RecordTest) TestBRecordWith1099SA(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099SaType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099SaJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099SaType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099SaAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099SaAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099SaAscii))
}

func (t *RecordTest) TestBRecordWith1099SB(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub1099SbType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord1099SbJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub1099SbType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099SbAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord1099SbAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord1099SbAscii))
}

func (t *RecordTest) TestBRecordWith3921(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub3921Type)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord3921Json, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub3921Type)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord3921Ascii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord3921Ascii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord3921Ascii))
}

func (t *RecordTest) TestBRecordWith3922(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub3922Type)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord3922Json, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub3922Type)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord3922Ascii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord3922Ascii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord3922Ascii))
}

func (t *RecordTest) TestBRecordWith5498(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub5498Type)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord5498Json, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub5498Type)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord5498Ascii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord5498Ascii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord5498Ascii))
}

func (t *RecordTest) TestBRecordWith5498ESA(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub5498EsaType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord5498EsaJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub5498EsaType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord5498EsaAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord5498EsaAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord5498EsaAscii))
}

func (t *RecordTest) TestBRecordWith5498SA(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.Sub5498SaType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecord5498SaJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.Sub5498SaType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord5498SaAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecord5498SaAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecord5498SaAscii))
}

func (t *RecordTest) TestBRecordWithW2G(c *check.C) {
	r := &BRecord{}
	err := r.SetTypeOfReturn(config.SubW2GType)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err = json.Unmarshal(t.bRecordW2GJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(r.extRecord.Type(), check.Equals, config.SubW2GType)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecordW2GAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.bRecordW2GAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.bRecordW2GAscii))
}
