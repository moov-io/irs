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
	person.Payees = append(person.Payees, records.NewBRecord("A"))
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
