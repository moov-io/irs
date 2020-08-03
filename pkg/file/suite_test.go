package file

import (
	"gopkg.in/check.v1"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type FileTest struct {
	oneTransactionJson     []byte
	oneTransactionAscii    []byte
	jsonWithInvalidPayment []byte
	jsonWithoutCRecord     []byte
	fileWithTestOptionJson []byte
}

var _ = check.Suite(&FileTest{})

func (t *FileTest) SetUpSuite(c *check.C) {
	var err error

	t.oneTransactionJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "oneTransactionFile.json"))
	c.Assert(err, check.IsNil)

	t.oneTransactionAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "oneTransactionFile.ascii"))
	c.Assert(err, check.IsNil)

	t.jsonWithInvalidPayment, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "fileWithInvalidPayment.json"))
	c.Assert(err, check.IsNil)

	t.jsonWithoutCRecord, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "fileWithoutCRecord.json"))
	c.Assert(err, check.IsNil)

	t.fileWithTestOptionJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "fileWithTestOption.json"))
	c.Assert(err, check.IsNil)
}
