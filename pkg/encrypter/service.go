// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package encrypter

type EncryptService interface {
	Encrypt(buf []byte) ([]byte, error)
	Decrypt(buf []byte) ([]byte, error)
	GetType() string
}

const (
	GCM        = "GCM"
	CBC        = "CBC"
	NonceSize  = 12
	EncryptKey = "Moov Irs Encryption AES-256 Key "
)
