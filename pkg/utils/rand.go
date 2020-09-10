// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

import (
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

const (
	letters       = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterBits    = 6
	letterIdxMask = 1<<letterBits - 1
	letterIdxMax  = 63 / letterBits
)

func RandAlphanumericString(length int) string {
	b := make([]byte, length)
	for i, cache, remain := length-1, source.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = source.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterBits
		remain--
	}
	return string(b)
}
