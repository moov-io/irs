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
}

func (t *FileTest) TestValidateWithOneTransactionJsonFile(c *check.C) {
	f, err := CreateFile(t.oneTransactionJson)
	c.Assert(err, check.IsNil)
	err = f.Validate()
	c.Assert(err, check.IsNil)
}
