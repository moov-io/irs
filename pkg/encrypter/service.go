// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package encrypter

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/moov-io/irs/pkg/utils"
)

type EncryptService interface {
	Encrypt(buf, non []byte) ([]byte, error)
	Decrypt(buf, non []byte) ([]byte, error)
	GetType() string
}

const (
	GCM          = "GCM"
	CBC          = "CBC"
	MinNonceSize = 12
	EncryptKey   = "Moov Irs Encryption AES-256 Key "
)

func NewEncryptService(key, method string) (EncryptService, error) {
	encrypt := &encryptInstance{}
	if len(key) == 0 {
		key = EncryptKey
	}
	encrypt.key = CreateKey(key)
	switch method {
	case GCM, CBC:
		encrypt.etype = method
	default:
		return nil, utils.ErrUnknownEncryptionType
	}

	return encrypt, nil
}

func GenerateNonce(id string, created time.Time) []byte {
	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	non := []byte(id + fmt.Sprintf("%d", created.Unix()))
	return []byte(hex.EncodeToString(non))
}

func CreateKey(key string) []byte {
	return []byte(hex.EncodeToString([]byte(key)))
}
