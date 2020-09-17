// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package documents

import (
	"database/sql"
	"io/ioutil"
	"path/filepath"
	"testing"

	"gopkg.in/check.v1"

	"github.com/moov-io/identity/pkg/database"
	"github.com/moov-io/irs/pkg/encrypter"
	"github.com/moov-io/irs/pkg/file"
	"github.com/moov-io/irs/pkg/utils"
)

func Test(t *testing.T) { check.TestingT(t) }

var _ = check.Suite(&DocumentTest{})

// Document test
type DocumentTest struct {
	db                 *sql.DB
	close              func()
	service            StorageService
	encrypt            encrypter.EncryptService
	oneTransactionJson []byte
}

func (t *DocumentTest) SetUpSuite(c *check.C) {
	var err error
	t.db, t.close, err = database.NewAndMigrate(database.InMemorySqliteConfig, nil, nil)
	c.Assert(err, check.IsNil)
	t.encrypt, err = encrypter.NewEncryptService("", encrypter.GCM)
	c.Assert(err, check.IsNil)
	t.service, err = NewStorageService(t.db, t.encrypt)
	c.Assert(err, check.IsNil)
	t.oneTransactionJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "oneTransactionFile.json"))
	c.Assert(err, check.IsNil)
}

func (t *DocumentTest) TearDownSuite(c *check.C) {
	defer t.close()
}

func (t *DocumentTest) TestDocumentCRUD(c *check.C) {
	f, err := file.CreateFile(t.oneTransactionJson)
	c.Assert(err, check.IsNil)
	id, err := utils.RandAlphanumericString(40)
	c.Assert(err, check.IsNil)
	doc := &DocumentInformation{File: f, DocumentID: id}
	err = t.service.Save(doc)
	c.Assert(err, check.IsNil)
	stored, err := t.service.Get(doc.DocumentID)
	c.Assert(err, check.IsNil)
	nonce := encrypter.GenerateNonce(stored.DocumentID, stored.Created.Time)
	ascii, err := t.encrypt.Decrypt(stored.Ascii, nonce)
	c.Assert(err, check.IsNil)
	c.Assert(f.Ascii(), check.DeepEquals, ascii)
	doc.DocumentID = ""
	err = t.service.Save(doc)
	c.Assert(err, check.IsNil)
}

func (t *DocumentTest) TestFailedDocumentSave(c *check.C) {
	f, err := file.CreateFile(t.oneTransactionJson)
	c.Assert(err, check.IsNil)
	id, err := utils.RandAlphanumericString(1)
	c.Assert(err, check.IsNil)
	doc := &DocumentInformation{File: f, DocumentID: id}
	err = t.service.Save(doc)
	c.Assert(err, check.NotNil)
	_, err = t.service.Get(doc.DocumentID)
	c.Assert(err, check.NotNil)
	old := documentSelect
	documentSelect = `id`
	_, err = t.service.Get(doc.DocumentID)
	c.Assert(err, check.NotNil)
	documentSelect = old
}
