// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"encoding/json"
	"gopkg.in/check.v1"

	"github.com/moov-io/irs/pkg/config"
)

func (t *RecordTest) TestARecord(c *check.C) {
	r := NewARecord()
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.aRecordJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.aRecordAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.aRecordAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.aRecordAscii))
	r.SetSequenceNumber(1)
	c.Assert(r.SequenceNumber(), check.Equals, 1)
	c.Assert(r.Validate(), check.IsNil)
	r.SetSequenceNumber(-1)
	c.Assert(r.Validate(), check.NotNil)
	c.Assert(r.Type(), check.Equals, config.ARecordType)
}

func (t *RecordTest) TestARecordWithError(c *check.C) {
	r := &ARecord{}
	err := r.Parse(t.aRecordAscii[1:])
	c.Assert(err, check.Not(check.IsNil))
	err = r.Parse(t.aRecordAscii)
	c.Assert(err, check.IsNil)
	r.CombinedFSFilingProgram = "ERR"
	c.Assert(r.ValidateCombinedFSFilingProgram(), check.Not(check.IsNil))
	r.LastFilingIndicator = "ERR"
	c.Assert(r.ValidateLastFilingIndicator(), check.Not(check.IsNil))
	r.TypeOfReturn = "ERR"
	c.Assert(r.ValidateTypeOfReturn(), check.Not(check.IsNil))
	r.ForeignEntityIndicator = "ERR"
	c.Assert(r.ValidateForeignEntityIndicator(), check.Not(check.IsNil))
	r.TransferAgentIndicator = "ERR"
	c.Assert(r.ValidateTransferAgentIndicator(), check.Not(check.IsNil))
	r.PayerState = "ERR"
	c.Assert(r.ValidatePayerState(), check.Not(check.IsNil))
	c.Assert(r.ValidateAmountCodes(), check.Not(check.IsNil))
}

func (t *RecordTest) TestARecord_BlankFields(c *check.C) {
	r := NewARecord()
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.aRecordJson, r)
	c.Assert(err, check.IsNil)
	transmitter := r.(*ARecord)
	transmitter.LastFilingIndicator = ""
	transmitter.ForeignEntityIndicator = ""
	c.Assert(r.Validate(), check.IsNil)
}
