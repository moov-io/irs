// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"testing"
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
