// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package encrypter

import (
	"testing"
	"time"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

var _ = check.Suite(&EncrypterTest{})

// Encrypter test
type EncrypterTest struct {
	plaintext []byte
}

func (t *EncrypterTest) SetUpSuite(c *check.C) {
	t.plaintext = []byte("exampleplaintext")
}

func (t *EncrypterTest) TestEncryptionWithGCM(c *check.C) {
	service, err := NewEncryptService(0, time.Now(), GCM)
	c.Assert(err, check.IsNil)
	encrypted, err := service.Encrypt(t.plaintext)
	c.Assert(err, check.IsNil)
	decrypted, err := service.Decrypt(encrypted)
	c.Assert(err, check.IsNil)
	c.Assert(t.plaintext, check.DeepEquals, decrypted)
}

func (t *EncrypterTest) TestEncryptionWithCBC(c *check.C) {
	service, err := NewEncryptService(0, time.Now(), CBC)
	c.Assert(err, check.IsNil)
	encrypted, err := service.Encrypt(t.plaintext)
	c.Assert(err, check.IsNil)
	decrypted, err := service.Decrypt(encrypted)
	c.Assert(err, check.IsNil)
	c.Assert(t.plaintext, check.DeepEquals, decrypted)
}

func (t *EncrypterTest) TestWithInvalidType(c *check.C) {
	_, err := NewEncryptService(0, time.Now(), "Unknown")
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "unsupported type")
	service := encryptInstance{Type: "Unknown"}
	_, err = service.Encrypt(t.plaintext)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "unsupported encrypt type")
	_, err = service.Decrypt(t.plaintext)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "unsupported decrypt type")
}

func (t *EncrypterTest) TestWithInvalidKey(c *check.C) {
	service := encryptInstance{
		Type: GCM,
		Key:  []byte("6368616e676520746869732070617373776f726420746f20612073656372657")}
	_, err := service.Encrypt(t.plaintext)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "encoding/hex: odd length hex string")
	_, err = service.Decrypt(t.plaintext)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "encoding/hex: odd length hex string")

	service.Key = []byte("6368616e676520746869732070617373776f726420746f2061207365637265")
	_, err = service.Encrypt(t.plaintext)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "crypto/aes: invalid key size 31")
	_, err = service.Decrypt(t.plaintext)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "crypto/aes: invalid key size 31")
	service.Type = CBC
	_, err = service.Encrypt(t.plaintext)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "crypto/aes: invalid key size 31")
	_, err = service.Decrypt(t.plaintext)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "crypto/aes: invalid key size 31")
}

func (t *EncrypterTest) TestWithInvalidNonce(c *check.C) {
	service := encryptInstance{
		Type:  GCM,
		Key:   []byte("6368616e676520746869732070617373776f726420746f206120736563726574"),
		Nonce: []byte("64a9433eae7ccceee2fc0ed")}
	_, err := service.Encrypt(t.plaintext)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "encoding/hex: odd length hex string")
	_, err = service.Decrypt(t.plaintext)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "encoding/hex: odd length hex string")
	service.Nonce = []byte("64a9433eae7ccceee2fc0e")
	_, err = service.Encrypt(t.plaintext)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "crypto/cipher: incorrect nonce length given to GCM")
	_, err = service.Decrypt(t.plaintext)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "crypto/cipher: incorrect nonce length given to GCM")
}

func (t *EncrypterTest) TestErrUsingCBC(c *check.C) {
	service := encryptInstance{
		Type:  CBC,
		Key:   []byte("6368616e676520746869732070617373776f726420746f206120736563726574"),
		Nonce: []byte("64a9433eae7ccceee2fc0eda")}
	_, err := service.Encrypt(t.plaintext[1:])
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "text is not a multiple of the block size")
	_, err = service.Decrypt(t.plaintext[:11])
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "text too short")
	_, err = service.Decrypt([]byte("exampleplaintextexampleplaintex"))
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "text is not a multiple of the block size")
}
