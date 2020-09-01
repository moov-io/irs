// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package documents

import (
	"github.com/moov-io/irs/pkg/file"
)

type Document struct {
	DocumentID string
	File       file.File
	Metadata   map[string]string
}

type Storage interface {
	Save(doc *Document) error
}
