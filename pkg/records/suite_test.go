// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/check.v1"

	"github.com/moov-io/irs/pkg/utils"
)

func Test(t *testing.T) { check.TestingT(t) }

type RecordTest struct {
	tRecordJson  []byte
	tRecordAscii []byte
	aRecordJson  []byte
	aRecordAscii []byte
	bRecordJson  []byte
	bRecordAscii []byte
}

var _ = check.Suite(&RecordTest{})

func (t *RecordTest) SetUpSuite(c *check.C) {
	f, err := os.Open(filepath.Join("..", "..", "test", "testdata", "transmitterRecord.json"))
	c.Assert(err, check.IsNil)
	t.tRecordJson = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "transmitterRecord.ascii"))
	c.Assert(err, check.IsNil)
	t.tRecordAscii = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "payerRecord.json"))
	c.Assert(err, check.IsNil)
	t.aRecordJson = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "payerRecord.ascii"))
	c.Assert(err, check.IsNil)
	t.aRecordAscii = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "payeeRecord.json"))
	c.Assert(err, check.IsNil)
	t.bRecordJson = utils.ReadFile(f)

	f, err = os.Open(filepath.Join("..", "..", "test", "testdata", "payeeRecord.ascii"))
	c.Assert(err, check.IsNil)
	t.bRecordAscii = utils.ReadFile(f)
}
