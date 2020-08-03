package file

import (
	"bytes"
	"encoding/json"
	"gopkg.in/check.v1"
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
