// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzzreader

import (
	"github.com/moov-io/irs/pkg/file"
)

// Return codes (from go-fuzz docs)
//
// The function must return 1 if the fuzzer should increase priority
// of the given input during subsequent fuzzing (for example, the input is
// lexically correct and was parsed successfully); -1 if the input must not be
// added to corpus even if gives new coverage; and 0 otherwise; other values are
// reserved for future use.
func Fuzz(data []byte) int {
	f, err := file.CreateFile(data)
	if err != nil {
		return 0
	}

	if _, err := f.TCC(); err != nil {
		return -1
	}

	if err := f.Validate(); err != nil {
		return 0
	}

	// Prioritize generated files with header, trailer, and data records.
	return 1
}
