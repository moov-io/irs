// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package documents

import (
	"database/sql"
)

// Document
type Document struct {
	DocumentID string       `json:"document_id"`
	Pdf        []byte       `json:"pdf,omitempty"`
	Ascii      []byte       `json:"ascii"`
	Created    sql.NullTime `json:"created_at,omitempty"`
	Deleted    sql.NullTime `json:"deleted_at"`
}
