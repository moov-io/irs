// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package service_test

import (
	"testing"

	"github.com/moov-io/base/config"
	"github.com/moov-io/base/log"
	"github.com/moov-io/irs/pkg/service"
	"github.com/stretchr/testify/require"
)

func Test_ConfigLoading(t *testing.T) {
	logger := log.NewNopLogger()

	ConfigService := config.NewService(logger)

	gc := &service.GlobalConfig{}
	err := ConfigService.Load(gc)
	require.Nil(t, err)
}
