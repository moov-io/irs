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
	nonce     []byte
}

func (t *EncrypterTest) SetUpSuite(c *check.C) {
	t.plaintext = []byte("exampleplaintext")
	t.nonce = GenerateNonce("test", time.Now())
}

func (t *EncrypterTest) TestEncryptionWithGCM(c *check.C) {
	service, err := NewEncryptService("", GCM)
	c.Assert(err, check.IsNil)
	encrypted, err := service.Encrypt(t.plaintext, t.nonce)
	c.Assert(err, check.IsNil)
	decrypted, err := service.Decrypt(encrypted, t.nonce)
	c.Assert(err, check.IsNil)
	c.Assert(t.plaintext, check.DeepEquals, decrypted)
}

func (t *EncrypterTest) TestEncryptionWithCBC(c *check.C) {
	service, err := NewEncryptService("", CBC)
	c.Assert(err, check.IsNil)
	encrypted, err := service.Encrypt(t.plaintext, t.nonce)
	c.Assert(err, check.IsNil)
	decrypted, err := service.Decrypt(encrypted, t.nonce)
	c.Assert(err, check.IsNil)
	c.Assert(t.plaintext, check.DeepEquals, decrypted)
}

func (t *EncrypterTest) TestWithInvalidType(c *check.C) {
	_, err := NewEncryptService("", "Unknown")
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "is unknown encryption type")
	service := encryptInstance{etype: "Unknown"}
	_, err = service.Encrypt(t.plaintext, t.nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "is unknown encryption type")
	_, err = service.Decrypt(t.plaintext, t.nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "is unknown encryption type")
}

func (t *EncrypterTest) TestWithInvalidKey(c *check.C) {
	service := encryptInstance{
		etype: GCM,
		key:   []byte("6368616e676520746869732070617373776f726420746f20612073656372657")}
	_, err := service.Encrypt(t.plaintext, t.nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "encoding/hex: odd length hex string")
	_, err = service.Decrypt(t.plaintext, t.nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "encoding/hex: odd length hex string")

	service.key = []byte("6368616e676520746869732070617373776f726420746f2061207365637265")
	_, err = service.Encrypt(t.plaintext, t.nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "crypto/aes: invalid key size 31")
	_, err = service.Decrypt(t.plaintext, t.nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "crypto/aes: invalid key size 31")
	service.etype = CBC
	_, err = service.Encrypt(t.plaintext, t.nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "crypto/aes: invalid key size 31")
	_, err = service.Decrypt(t.plaintext, t.nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "crypto/aes: invalid key size 31")
}

func (t *EncrypterTest) TestWithInvalidNonce(c *check.C) {
	service := encryptInstance{
		etype: GCM,
		key:   []byte("6368616e676520746869732070617373776f726420746f206120736563726574")}
	nonce := []byte("64a9433eae7ccceee2fc0ed")
	_, err := service.Encrypt(t.plaintext, nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "encoding/hex: odd length hex string")
	_, err = service.Decrypt(t.plaintext, nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "encoding/hex: odd length hex string")
	nonce = []byte("64a9433eae7ccceee2fc0e")
	_, err = service.Encrypt(t.plaintext, nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "crypto/cipher: incorrect nonce length given to GCM")
	_, err = service.Decrypt(t.plaintext, nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "crypto/cipher: incorrect nonce length given to GCM")
}

func (t *EncrypterTest) TestErrUsingCBC(c *check.C) {
	service := encryptInstance{
		etype: CBC,
		key:   []byte("6368616e676520746869732070617373776f726420746f206120736563726574")}
	nonce := []byte("64a9433eae7ccceee2fc0eda")
	_, err := service.Encrypt(t.plaintext[1:], nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "text is not a multiple of the block size")
	_, err = service.Decrypt(t.plaintext[:11], nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "text too short")
	_, err = service.Decrypt([]byte("exampleplaintextexampleplaintex"), nonce)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "text is not a multiple of the block size")
}
