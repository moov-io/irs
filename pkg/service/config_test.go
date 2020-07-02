package service_test

import (
	"testing"

	"github.com/moov-io/identity/pkg/config"
	"github.com/moov-io/identity/pkg/logging"
	"github.com/stretchr/testify/require"
	"irs/pkg/service"
)

func Test_ConfigLoading(t *testing.T) {
	logger := logging.NewNopLogger()

	ConfigService := config.NewConfigService(logger)

	gc := &service.GlobalConfig{}
	err := ConfigService.Load(gc)
	require.Nil(t, err)
}
