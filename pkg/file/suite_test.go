// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"gopkg.in/check.v1"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type FileTest struct {
	oneTransactionJson                 []byte
	oneTransactionAscii                []byte
	jsonWithInvalidPayment             []byte
	jsonWithoutCRecord                 []byte
	fileWithTestOptionJson             []byte
	oneTransactionWithoutKJson         []byte
	oneTransactionFileInvalidStateJson []byte
	sample1099IntJson                  []byte
	sample1099MiscJson                 []byte
	sample1099OidJson                  []byte
	sample1099PatrJson                 []byte
}

var _ = check.Suite(&FileTest{})

func (t *FileTest) SetUpSuite(c *check.C) {
	var err error

	t.oneTransactionJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "oneTransactionFile.json"))
	c.Assert(err, check.IsNil)

	t.oneTransactionWithoutKJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "oneTransactionFileWithoutK.json"))
	c.Assert(err, check.IsNil)

	t.oneTransactionFileInvalidStateJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "oneTransactionFileInvalidState.json"))
	c.Assert(err, check.IsNil)

	t.oneTransactionAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "oneTransactionFile.ascii"))
	c.Assert(err, check.IsNil)

	t.jsonWithInvalidPayment, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "fileWithInvalidPayment.json"))
	c.Assert(err, check.IsNil)

	t.jsonWithoutCRecord, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "fileWithoutCRecord.json"))
	c.Assert(err, check.IsNil)

	t.fileWithTestOptionJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "fileWithTestOption.json"))
	c.Assert(err, check.IsNil)

	t.sample1099IntJson, err = ioutil.ReadFile(filepath.Join("..", "..", "docs", "examples", "1099int.json"))
	c.Assert(err, check.IsNil)

	t.sample1099MiscJson, err = ioutil.ReadFile(filepath.Join("..", "..", "docs", "examples", "1099misc.json"))
	c.Assert(err, check.IsNil)

	t.sample1099OidJson, err = ioutil.ReadFile(filepath.Join("..", "..", "docs", "examples", "1099oid.json"))
	c.Assert(err, check.IsNil)

	t.sample1099PatrJson, err = ioutil.ReadFile(filepath.Join("..", "..", "docs", "examples", "1099patr.json"))
	c.Assert(err, check.IsNil)
}
