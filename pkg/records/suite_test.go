// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type RecordTest struct {
	tRecordJson          []byte
	tRecordAscii         []byte
	aRecordJson          []byte
	aRecordAscii         []byte
	bRecord1099MiscJson  []byte
	bRecord1099MiscAscii []byte
	bRecord1099NecJson   []byte
	bRecord1099NecAscii  []byte
	bRecord1099IntJson   []byte
	bRecord1099IntAscii  []byte
	bRecord1099OidJson   []byte
	bRecord1099OidAscii  []byte
	bRecord1099PatrJson  []byte
	bRecord1099PatrAscii []byte
	bRecord1097BtcJson   []byte
	bRecord1097BtcAscii  []byte
	bRecord1098Json      []byte
	bRecord1098Ascii     []byte
	bRecord1098CJson     []byte
	bRecord1098CAscii    []byte
	bRecord1098EJson     []byte
	bRecord1098EAscii    []byte
	bRecord1098FJson     []byte
	bRecord1098FAscii    []byte
	bRecord1098QJson     []byte
	bRecord1098QAscii    []byte
	bRecord1098TJson     []byte
	bRecord1098TAscii    []byte
	bRecord1099AJson     []byte
	bRecord1099AAscii    []byte
	bRecord1099BJson     []byte
	bRecord1099BAscii    []byte
	bRecord1099CJson     []byte
	bRecord1099CAscii    []byte
	bRecord1099CapJson   []byte
	bRecord1099CapAscii  []byte
	bRecord1099DivJson   []byte
	bRecord1099DivAscii  []byte
	bRecord1099GJson     []byte
	bRecord1099GAscii    []byte
	bRecord1099HJson     []byte
	bRecord1099HAscii    []byte
	bRecord1099KJson     []byte
	bRecord1099KAscii    []byte
	bRecord1099LsJson    []byte
	bRecord1099LsAscii   []byte
	bRecord1099LtcJson   []byte
	bRecord1099LtcAscii  []byte
	bRecord1099QJson     []byte
	bRecord1099QAscii    []byte
	bRecord1099RJson     []byte
	bRecord1099RAscii    []byte
	bRecord1099SJson     []byte
	bRecord1099SAscii    []byte
	bRecord1099SaJson    []byte
	bRecord1099SaAscii   []byte
	bRecord1099SbJson    []byte
	bRecord1099SbAscii   []byte
	bRecord3921Json      []byte
	bRecord3921Ascii     []byte
	bRecord3922Json      []byte
	bRecord3922Ascii     []byte
	bRecord5498Json      []byte
	bRecord5498Ascii     []byte
	bRecord5498EsaJson   []byte
	bRecord5498EsaAscii  []byte
	bRecord5498SaJson    []byte
	bRecord5498SaAscii   []byte
	bRecordW2GJson       []byte
	bRecordW2GAscii      []byte
	cRecordJson          []byte
	cRecordAscii         []byte
	kRecordJson          []byte
	kRecordAscii         []byte
	fRecordJson          []byte
	fRecordAscii         []byte
}

var _ = check.Suite(&RecordTest{})

func (t *RecordTest) SetUpSuite(c *check.C) {
	var err error

	t.tRecordJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "transmitterRecord.json"))
	c.Assert(err, check.IsNil)
	t.tRecordAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "transmitterRecord.ascii"))
	c.Assert(err, check.IsNil)

	t.aRecordJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payerRecord.json"))
	c.Assert(err, check.IsNil)
	t.aRecordAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payerRecord.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099MiscJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Misc.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099MiscAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Misc.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099NecJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Nec.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099NecAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Nec.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099IntJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Int.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099IntAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Int.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099OidJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Oid.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099OidAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Oid.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099PatrJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Patr.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099PatrAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Patr.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1097BtcJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1097Btc.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1097BtcAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1097Btc.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1098Json, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1098.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1098Ascii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1098.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1098CJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1098C.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1098CAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1098C.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1098EJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1098E.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1098EAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1098E.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1098FJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1098F.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1098FAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1098F.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1098QJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1098Q.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1098QAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1098Q.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1098TJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1098T.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1098TAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1098T.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099AJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099A.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099AAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099A.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099BJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099B.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099BAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099B.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099CJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099C.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099CAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099C.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099CapJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Cap.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099CapAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Cap.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099DivJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Div.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099DivAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Div.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099GJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099G.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099GAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099G.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099HJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099H.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099HAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099H.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099KJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099K.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099KAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099K.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099LsJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Ls.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099LsAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Ls.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099LtcJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Ltc.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099LtcAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Ltc.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099QJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Q.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099QAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Q.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099RJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099R.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099RAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099R.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099SJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099S.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099SAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099S.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099SaJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Sa.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099SaAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Sa.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099SbJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Sb.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099SbAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Sb.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099SbJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Sb.json"))
	c.Assert(err, check.IsNil)
	t.bRecord1099SbAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Sb.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord3921Json, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith3921.json"))
	c.Assert(err, check.IsNil)
	t.bRecord3921Ascii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith3921.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord3922Json, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith3922.json"))
	c.Assert(err, check.IsNil)
	t.bRecord3922Ascii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith3922.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord5498Json, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith5498.json"))
	c.Assert(err, check.IsNil)
	t.bRecord5498Ascii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith5498.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord5498EsaJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith5498Esa.json"))
	c.Assert(err, check.IsNil)
	t.bRecord5498EsaAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith5498Esa.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord5498SaJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith5498Sa.json"))
	c.Assert(err, check.IsNil)
	t.bRecord5498SaAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith5498Sa.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecordW2GJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWithW2G.json"))
	c.Assert(err, check.IsNil)
	t.bRecordW2GAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWithW2G.ascii"))
	c.Assert(err, check.IsNil)

	t.cRecordJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "endPayerRecord.json"))
	c.Assert(err, check.IsNil)
	t.cRecordAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "endPayerRecord.ascii"))
	c.Assert(err, check.IsNil)

	t.kRecordJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "stateRecord.json"))
	c.Assert(err, check.IsNil)
	t.kRecordAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "stateRecord.ascii"))
	c.Assert(err, check.IsNil)

	t.fRecordJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "endTransmitterRecord.json"))
	c.Assert(err, check.IsNil)
	t.fRecordAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "endTransmitterRecord.ascii"))
	c.Assert(err, check.IsNil)
}
