// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"strings"
	"testing"

	PDF "github.com/moov-io/irs/pkg/pdf_generator"
	"github.com/moov-io/irs/pkg/records"
)

func TestEQ(t *testing.T) {
	m1 := make(map[string]bool)
	m2 := make(map[string]bool)

	m1["foo"] = false
	if eq(m1, m2) {
		t.Error("expected not equal")
	}

	m2["foo"] = true
	if eq(m1, m2) {
		t.Error("expected not equal")
	}

	m2["foo"] = false
	if !eq(m1, m2) {
		t.Error("expected equal")
	}
}

func TestFillAmounts(t *testing.T) {
	pdf := &PDF.Pdf1099Misc{}
	amountCodes := strings.Split("123456789ABCDEFG", "")
	record := &records.CRecord{}
	err := fillAmounts(amountCodes, pdf, record)
	if err != nil {
		t.Errorf(err.Error())
	}
}
