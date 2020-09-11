// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package documents

import (
	"database/sql"
	"fmt"
	"time"

	encrypt "github.com/moov-io/irs/pkg/encrypter"
	"github.com/moov-io/irs/pkg/file"
	"github.com/moov-io/irs/pkg/utils"
)

// DocumentInformation
type DocumentInformation struct {
	DocumentID string
	File       file.File
	Metadata   map[string]string
}

// StorageService
type StorageService interface {
	Save(doc *DocumentInformation) error
	Get(id string) (*Document, error)
}

// NewStorageService
func NewStorageService(db *sql.DB, encrypter encrypt.EncryptService) (StorageService, error) {
	return &storageService{db: db, encrypter: encrypter}, nil
}

type storageService struct {
	db        *sql.DB
	encrypter encrypt.EncryptService
}

var documentSelect = `
	document_id, 
	pdf, 
	ascii, 
	created_at, 
	deleted_at
`

func (s *storageService) Save(doc *DocumentInformation) error {
	if doc == nil && doc.File == nil {
		return utils.ErrNullFile
	}

	var err error
	id := doc.DocumentID
	if len(id) == 0 {
		id, err = utils.RandAlphanumericString(40)
		if err != nil {
			return err
		}
	}

	ascii := doc.File.Ascii()
	created := time.Now()
	if s.encrypter != nil {
		nonce := encrypt.GenerateNonce(doc.DocumentID, created)
		ascii, err = s.encrypter.Encrypt(ascii, nonce)
		if err != nil {
			return err
		}
	}

	qry := `
		INSERT INTO documents(
			document_id, 
			pdf,
			ascii, 
			created_at
		) VALUES (?, ?, ?, ?)
	`
	res, err := s.db.Exec(qry,
		id,
		nil,
		ascii,
		created)

	if err != nil {
		return err
	}

	if cnt, err := res.RowsAffected(); cnt != 1 || err != nil {
		return sql.ErrNoRows
	}
	return nil
}

func (s *storageService) Get(id string) (*Document, error) {
	qry := fmt.Sprintf(`
		SELECT %s
		FROM documents
		WHERE document_id = ?
		LIMIT 1
	`, documentSelect)

	results, err := s.queryScan(qry, id)
	if err != nil {
		return nil, err
	}

	if len(results) != 1 {
		return nil, sql.ErrNoRows
	}

	return &results[0], nil
}

func (s *storageService) queryScan(query string, args ...interface{}) ([]Document, error) {
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []Document
	for rows.Next() {
		element := Document{}
		if err := rows.Scan(&element.DocumentID, &element.Pdf, &element.Ascii, &element.Created, &element.Deleted); err != nil {
			return nil, err
		}
		documents = append(documents, element)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return documents, nil
}
