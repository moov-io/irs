// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"encoding/json"
	"github.com/moov-io/irs/pkg/records"
	"gopkg.in/check.v1"
	"strings"
)

func (t *FileTest) TestParseWithOneTransactionJsonFile(c *check.C) {
	f1, err := CreateFile(t.oneTransactionJson)
	c.Assert(err, check.IsNil)
	buf1, err := json.Marshal(f1)
	c.Assert(err, check.IsNil)
	var prettyJSON1 bytes.Buffer
	json.Indent(&prettyJSON1, buf1, "", "  ")
	ascii := f1.Ascii()
	c.Assert(string(ascii), check.Equals, string(t.oneTransactionAscii))
	_, err = f1.Pdf()
	c.Assert(err, check.IsNil)
	f2, err := CreateFile(ascii)
	c.Assert(err, check.IsNil)
	buf2, err := json.Marshal(f2)
	c.Assert(err, check.IsNil)
	var prettyJSON2 bytes.Buffer
	json.Indent(&prettyJSON2, buf2, "", "  ")
	c.Assert(prettyJSON1.String(), check.Equals, prettyJSON2.String())
	err = f1.Validate()
	c.Assert(err, check.IsNil)
	err = f2.Validate()
	c.Assert(err, check.IsNil)
	tcc, err := f1.TCC()
	c.Assert(err, check.IsNil)
	c.Assert(tcc, check.NotNil)
	c.Assert(f1.SetTCC("123456"), check.NotNil)
	c.Assert(f1.SetTCC("12345"), check.IsNil)
	p := &paymentPerson{}
	c.Assert(p.Type(), check.Equals, "Person")
}

func (t *FileTest) TestValidateWithOneTransactionJsonFile(c *check.C) {
	f := &fileInstance{}
	err := json.Unmarshal(t.oneTransactionJson, f)
	c.Assert(err, check.IsNil)
	err = f.Validate()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestParseFailed(c *check.C) {
	f := &fileInstance{}
	err := json.Unmarshal(t.oneTransactionAscii, f)
	c.Assert(err, check.NotNil)
	err = f.Parse(t.oneTransactionJson)
	c.Assert(err, check.NotNil)
	r := records.NewARecord()
	err = readJsonWithRecord(r, t.fileWithTestOptionJson)
	c.Assert(err, check.NotNil)
	p := &paymentPerson{}
	err = readJsonWithPerson(p, t.fileWithTestOptionJson)
	c.Assert(err, check.NotNil)

	f1, err := CreateFile(t.oneTransactionJson)
	c.Assert(err, check.IsNil)
	err = f1.Validate()
	c.Assert(err, check.IsNil)
	file1, ok := f1.(*fileInstance)
	c.Assert(ok, check.Equals, true)
	person := file1.PaymentPersons[0]
	err = person.Validate()
	c.Assert(err, check.IsNil)
	person = file1.PaymentPersons[0]
	person.States = append(person.States, records.NewKRecord())
	c.Assert(person.Validate(), check.NotNil)
	person.EndPayer = records.NewCRecord()
	c.Assert(person.Validate(), check.NotNil)
	_r, err := records.NewBRecord("1099-MISC")
	c.Assert(err, check.IsNil)
	person.Payees = append(person.Payees, _r)
	c.Assert(person.Validate(), check.NotNil)
	person.SetSequenceNumber(0)
	person.Payer = nil
	c.Assert(person.SequenceNumber(), check.Equals, 0)

	_, err = person.Parse(t.oneTransactionAscii)
	c.Assert(err, check.NotNil)
	_, err = person.Parse(t.oneTransactionAscii[750 : 750+749])
	c.Assert(err, check.NotNil)
	payerStr := string(t.oneTransactionAscii[750 : 750+750])
	payerStr = strings.ReplaceAll(payerStr, "ASDF1A 7", "ASDF1_ 7")
	_, err = person.Parse([]byte(payerStr))
	c.Assert(err, check.NotNil)
	payerStr = strings.ReplaceAll(payerStr, "A20171", "A_0171")
	_, err = person.Parse([]byte(payerStr))
	c.Assert(err, check.NotNil)
	_, err = person.Parse(t.oneTransactionAscii[750 : 750+750+749])
	c.Assert(err, check.NotNil)
	_, err = person.Parse(t.oneTransactionAscii[750 : 750+750+750])
	c.Assert(err, check.NotNil)
	payeeStr := string(t.oneTransactionAscii[750 : 750+750+750])
	payeeStr = strings.ReplaceAll(payeeStr, "B2017", "B_017")
	_, err = person.Parse([]byte(payeeStr))
	c.Assert(err, check.NotNil)
	endPayerStr := string(t.oneTransactionAscii[750:])
	endPayerStr = strings.ReplaceAll(endPayerStr, "K00000002", "K        ")
	_, err = person.Parse([]byte(endPayerStr))
	c.Assert(err, check.NotNil)
	endPayerStr = strings.ReplaceAll(endPayerStr, "C00000002", "C        ")
	_, err = person.Parse([]byte(endPayerStr))
	c.Assert(err, check.NotNil)
	endPayerStr = strings.ReplaceAll(endPayerStr, "C2017", "C_017")
	_, err = person.Parse([]byte(endPayerStr))
	c.Assert(err, check.NotNil)

	person = &paymentPerson{}
	c.Assert(person.validateRecords(), check.NotNil)
	c.Assert(person.validatePaymentCodes(), check.NotNil)
	c.Assert(person.integrationCheck(), check.NotNil)
	_, err = person.getTypeOfReturn()
	c.Assert(err, check.NotNil)
	person.Payer = records.NewKRecord()
	_, err = person.getTypeOfReturn()
	c.Assert(err, check.NotNil)
	_, _, err = person.getRecords()
	c.Assert(err, check.NotNil)
	person.Payer = records.NewARecord()
	_, err = person.getTypeOfReturn()
	c.Assert(err, check.NotNil)
	person.EndPayer = records.NewKRecord()
	c.Assert(person.validateRecords(), check.NotNil)
	_, _, err = person.getRecords()
	c.Assert(err, check.NotNil)
}

func (t *FileTest) TestFileInstanceErrorCases(c *check.C) {
	instance := &fileInstance{}
	_, _, err := instance.getRecords()
	c.Assert(err, check.NotNil)
	err = instance.integrationCheck()
	c.Assert(err, check.NotNil)
	err = instance.validateRecords()
	c.Assert(err, check.NotNil)
	err = instance.SetTCC("test-tcc")
	c.Assert(err, check.NotNil)
	_, err = instance.TCC()
	c.Assert(err, check.NotNil)
	instance.Transmitter = records.NewARecord()
	instance.EndTransmitter = records.NewARecord()
	_, _, err = instance.getRecords()
	c.Assert(err, check.NotNil)
	err = instance.validateRecords()
	c.Assert(err, check.NotNil)
	err = instance.validateRecordSequenceNumber()
	c.Assert(err, check.NotNil)
	instance.Transmitter = records.NewTRecord()
	_, _, err = instance.getRecords()
	c.Assert(err, check.NotNil)

	instance.Transmitter = nil
	instance.EndTransmitter = nil
	err = instance.Parse(t.oneTransactionAscii)
	c.Assert(err, check.IsNil)
	instance.EndTransmitter.SetSequenceNumber(1)
	err = instance.Validate()
	c.Assert(err, check.NotNil)
	err = instance.Validate()
	c.Assert(err, check.NotNil)
	err = instance.validateRecordSequenceNumber()
	c.Assert(err, check.NotNil)
	instance.PaymentPersons[0].States[0].SetSequenceNumber(0)
	err = instance.validateRecordSequenceNumber()
	c.Assert(err, check.NotNil)
	instance.PaymentPersons[0].EndPayer.SetSequenceNumber(0)
	err = instance.validateRecordSequenceNumber()
	c.Assert(err, check.NotNil)
	instance.PaymentPersons[0].Payees[0].SetSequenceNumber(0)
	err = instance.validateRecordSequenceNumber()
	c.Assert(err, check.NotNil)
	instance.PaymentPersons[0].Payer.SetSequenceNumber(0)
	err = instance.validateRecordSequenceNumber()
	c.Assert(err, check.NotNil)
	instance.Transmitter.SetSequenceNumber(0)
	err = instance.validateRecordSequenceNumber()
	c.Assert(err, check.NotNil)
	fRecord, ok := instance.EndTransmitter.(*records.FRecord)
	c.Assert(ok, check.Equals, true)
	tRecord, ok := instance.Transmitter.(*records.TRecord)
	c.Assert(ok, check.Equals, true)
	fRecord.TotalNumberPayees = tRecord.TotalNumberPayees - 1
	tRecord.TotalNumberPayees = 0
	err = instance.integrationCheck()
	c.Assert(err, check.NotNil)
	fRecord.NumberPayerRecords -= 1
	err = instance.integrationCheck()
	c.Assert(err, check.NotNil)
	a, _ := instance.PaymentPersons[0].Payer.(*records.ARecord)
	a.TypeOfReturn = "U"
	_, err = instance.Pdf()
	c.Assert(err, check.NotNil)
	instance.PaymentPersons[0].Payer = records.NewCRecord()
	_, err = instance.Pdf()
	c.Assert(err, check.NotNil)
	instance.PaymentPersons[0].Payer = nil
	_, err = instance.Pdf()
	c.Assert(err, check.NotNil)
}

func (t *FileTest) TestPersonErrorCases(c *check.C) {
	f1, err := CreateFile(t.oneTransactionJson)
	c.Assert(err, check.IsNil)
	file1, ok := f1.(*fileInstance)
	c.Assert(ok, check.Equals, true)
	person := file1.PaymentPersons[0]
	err = person.Validate()
	c.Assert(err, check.IsNil)
	kRecord, ok := person.States[0].(*records.KRecord)
	c.Assert(ok, check.Equals, true)
	kRecord.ControlTotal7 = 1
	c.Assert(person.validatePaymentCodes(), check.NotNil)
	kRecord.ControlTotal8 = 1
	c.Assert(person.validatePaymentCodes(), check.NotNil)
	person.States = append(person.States, records.NewCRecord())
	c.Assert(person.validatePaymentCodes(), check.NotNil)
	bRecord, ok := person.Payees[0].(*records.BRecord)
	c.Assert(ok, check.Equals, true)
	bRecord.PaymentAmount2 = 1
	c.Assert(person.validatePaymentCodes(), check.NotNil)
	person.Payees = append(person.Payees, records.NewCRecord())
	c.Assert(person.validatePaymentCodes(), check.NotNil)
	c.Assert(person.integrationCheck(), check.NotNil)
}

func (t *FileTest) TestFileWithInvalidPayment(c *check.C) {
	f, err := CreateFile(t.jsonWithInvalidPayment)
	c.Assert(err, check.IsNil)
	err = f.Validate()
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "has unexpected totals of any payment amount")
}

func (t *FileTest) TestFileWithoutCRecord(c *check.C) {
	f, err := CreateFile(t.jsonWithoutCRecord)
	c.Assert(err, check.IsNil)
	err = f.Validate()
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "should exist end of payer record")
}

func (t *FileTest) TestFileWithTestOption(c *check.C) {
	f, err := CreateFile(t.fileWithTestOptionJson)
	c.Assert(err, check.IsNil)
	err = f.Validate()
	c.Assert(err, check.NotNil)
}

func (t *FileTest) TestSample1099IntJson(c *check.C) {
	f1, err := CreateFile(t.sample1099IntJson)
	c.Assert(err, check.IsNil)
	buf1, err := json.Marshal(f1)
	c.Assert(err, check.IsNil)
	var prettyJSON1 bytes.Buffer
	json.Indent(&prettyJSON1, buf1, "", "  ")
	ascii := f1.Ascii()
	f2, err := CreateFile(ascii)
	c.Assert(err, check.IsNil)
	buf2, err := json.Marshal(f2)
	c.Assert(err, check.IsNil)
	var prettyJSON2 bytes.Buffer
	json.Indent(&prettyJSON2, buf2, "", "  ")
	c.Assert(prettyJSON1.String(), check.Equals, prettyJSON2.String())
	err = f1.Validate()
	c.Assert(err, check.IsNil)
	err = f2.Validate()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestSample1099OidJson(c *check.C) {
	f1, err := CreateFile(t.sample1099OidJson)
	c.Assert(err, check.IsNil)
	buf1, err := json.Marshal(f1)
	c.Assert(err, check.IsNil)
	var prettyJSON1 bytes.Buffer
	json.Indent(&prettyJSON1, buf1, "", "  ")
	ascii := f1.Ascii()
	f2, err := CreateFile(ascii)
	c.Assert(err, check.IsNil)
	buf2, err := json.Marshal(f2)
	c.Assert(err, check.IsNil)
	var prettyJSON2 bytes.Buffer
	json.Indent(&prettyJSON2, buf2, "", "  ")
	c.Assert(prettyJSON1.String(), check.Equals, prettyJSON2.String())
	err = f1.Validate()
	c.Assert(err, check.IsNil)
	err = f2.Validate()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestSample1099MiscJson(c *check.C) {
	f1, err := CreateFile(t.sample1099MiscJson)
	c.Assert(err, check.IsNil)
	buf1, err := json.Marshal(f1)
	c.Assert(err, check.IsNil)
	var prettyJSON1 bytes.Buffer
	json.Indent(&prettyJSON1, buf1, "", "  ")
	ascii := f1.Ascii()
	f2, err := CreateFile(ascii)
	c.Assert(err, check.IsNil)
	buf2, err := json.Marshal(f2)
	c.Assert(err, check.IsNil)
	var prettyJSON2 bytes.Buffer
	json.Indent(&prettyJSON2, buf2, "", "  ")
	c.Assert(prettyJSON1.String(), check.Equals, prettyJSON2.String())
	err = f1.Validate()
	c.Assert(err, check.IsNil)
	err = f2.Validate()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestSample1099PatrJson(c *check.C) {
	f1, err := CreateFile(t.sample1099PatrJson)
	c.Assert(err, check.IsNil)
	buf1, err := json.Marshal(f1)
	c.Assert(err, check.IsNil)
	var prettyJSON1 bytes.Buffer
	json.Indent(&prettyJSON1, buf1, "", "  ")
	ascii := f1.Ascii()
	f2, err := CreateFile(ascii)
	c.Assert(err, check.IsNil)
	buf2, err := json.Marshal(f2)
	c.Assert(err, check.IsNil)
	var prettyJSON2 bytes.Buffer
	json.Indent(&prettyJSON2, buf2, "", "  ")
	c.Assert(prettyJSON1.String(), check.Equals, prettyJSON2.String())
	err = f1.Validate()
	c.Assert(err, check.IsNil)
	err = f2.Validate()
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestOneTransactionFileWithoutKJson(c *check.C) {
	f1, err := CreateFile(t.oneTransactionWithoutKJson)
	c.Assert(err, check.IsNil)
	err = f1.Validate()
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "should be payee B records and the state totals K records")
}

func (t *FileTest) TestOneTransactionFileInvalidStateJson(c *check.C) {
	f1, err := CreateFile(t.oneTransactionFileInvalidStateJson)
	c.Assert(err, check.IsNil)
	err = f1.Validate()
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "is invalid combined federal/tate code in K record")
}
