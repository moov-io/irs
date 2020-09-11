// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"

	"github.com/moov-io/irs/pkg/utils"
)

var _ EncryptService = &encryptInstance{}

type encryptInstance struct {
	key   []byte
	etype string
}

func (e *encryptInstance) Encrypt(buf, nonce []byte) ([]byte, error) {
	strKey, err := hex.DecodeString(string(e.key))
	if err != nil {
		return nil, err
	}
	switch e.etype {
	case GCM:
		strNonce, err := hex.DecodeString(string(nonce))
		if err != nil {
			return nil, err
		}
		return gcmEncrypt(strKey, strNonce, buf)
	case CBC:
		return cbcEncrypt(strKey, buf)
	}
	return nil, utils.ErrUnknownEncryptionType
}

func (e *encryptInstance) Decrypt(buf, nonce []byte) ([]byte, error) {
	strKey, err := hex.DecodeString(string(e.key))
	if err != nil {
		return nil, err
	}
	switch e.etype {
	case GCM:
		strNonce, err := hex.DecodeString(string(nonce))
		if err != nil {
			return nil, err
		}
		return gcmDecrypt(strKey, strNonce, buf)
	case CBC:
		return cbcDecrypt(strKey, buf)
	}
	return nil, utils.ErrUnknownEncryptionType
}

func (e *encryptInstance) GetType() string {
	return e.etype
}

func gcmEncrypt(key, nonce, buf []byte) ([]byte, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(nonce) < MinNonceSize {
		return nil, utils.ErrInvalidNonceLength
	}

	aesGcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return nil, err
	}

	return aesGcm.Seal(nil, nonce, buf, nil), nil
}

func gcmDecrypt(key, nonce, buf []byte) ([]byte, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(nonce) < MinNonceSize {
		return nil, utils.ErrInvalidNonceLength
	}

	aesGcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return nil, err
	}

	return aesGcm.Open(nil, nonce, buf, nil)
}

func cbcEncrypt(key, buf []byte) ([]byte, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(buf)%aes.BlockSize != 0 {
		return nil, errors.New("text is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(buf))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], buf)

	return ciphertext, nil
}

func cbcDecrypt(key, buf []byte) ([]byte, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := buf
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("text too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("text is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext, nil
}
