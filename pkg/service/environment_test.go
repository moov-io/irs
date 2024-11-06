// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package service_test

import (
	"os"
	"runtime"
	"testing"

	"github.com/moov-io/base/database"
	logging "github.com/moov-io/base/log"
	"github.com/moov-io/irs/pkg/service"

	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
)

func Test_Environment_Startup(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") != "" {
		if runtime.GOOS != "linux" {
			t.Skip("Docker doens't work outside of linux on Actions")
		}
	}

	a := assert.New(t)

	env := &service.Environment{
		Logger: logging.NewLogger(log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))),
		Config: &service.Config{
			Database: database.DatabaseConfig{
				DatabaseName: "irs",
				MySQL: &database.MySQLConfig{
					Address:  "tcp(localhost:3306)",
					User:     "irs",
					Password: "irs",
				},
			},
		},
	}

	env, err := service.NewEnvironment(env)
	a.Nil(err)

	shutdown := env.RunServers(false)
	t.Cleanup(shutdown)
}
