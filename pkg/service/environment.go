// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"database/sql"
	"io/ioutil"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/moov-io/base/config"
	"github.com/moov-io/base/database"
	logging "github.com/moov-io/base/log"
	tmwLogging "github.com/moov-io/identity/pkg/logging"
	"github.com/moov-io/identity/pkg/stime"
	tmw "github.com/moov-io/tumbler/pkg/middleware"
	"github.com/moov-io/tumbler/pkg/webkeys"
)

// Environment - Contains everything that has been instantiated for this service.
type Environment struct {
	Logger        logging.Logger
	TumblerLogger tmwLogging.Logger
	Config        *Config
	TimeService   *stime.TimeService
	GatewayKeys   webkeys.WebKeysService
	PublicRouter  *mux.Router
	Shutdown      func()
}

// NewEnvironment - Generates a new default environment. Overrides can be specified via configs.
func NewEnvironment(env *Environment) (*Environment, error) {
	if env == nil {
		env = &Environment{}
	}

	if env.Logger == nil {
		env.Logger = logging.NewDefaultLogger()
	}

	if env.TumblerLogger == nil {
		env.TumblerLogger = tmwLogging.NewDefaultLogger()
	}

	if env.Config == nil {
		ConfigService := config.NewService(env.Logger)

		global := &GlobalConfig{}
		if err := ConfigService.Load(&global); err != nil {
			return nil, err
		}

		env.Config = &global.IRS
	}

	//db setup
	db, shutdownFn, err := initializeDatabase(env.Logger, env.Config.Database)
	if err != nil {
		shutdownFn()
		return nil, err
	}
	_ = db // delete once used.

	if env.TimeService == nil {
		t := stime.NewSystemTimeService()
		env.TimeService = &t
	}

	// router
	if env.PublicRouter == nil {
		env.PublicRouter = mux.NewRouter()
	}

	// auth middleware for the tokens coming from the gateway
	GatewayMiddleware, err := tmw.NewTumblerMiddlewareFromConfig(env.TumblerLogger, *env.TimeService, env.Config.Gateway)
	if err != nil {
		return nil, env.Logger.Fatal().LogErrorf("unable to startup the gateway middleware - %w", err).Err()
	}

	GatewayRouter := env.PublicRouter.NewRoute().Subrouter()
	GatewayRouter.Use(GatewayMiddleware.Handler)

	// configure custom handlers
	err = ConfigureHandlers(env.PublicRouter)
	if err != nil {
		return nil, env.Logger.LogErrorf("failed to configure handlers: %v", err).Err()
	}

	env.Shutdown = func() {
		shutdownFn()
	}

	return env, nil
}

func initializeDatabase(logger logging.Logger, config database.DatabaseConfig) (*sql.DB, func(), error) {
	ctx, cancelFunc := context.WithCancel(context.Background())

	// migrate database
	db, err := database.New(ctx, logger, config)
	if err != nil {
		return nil, cancelFunc, logger.Fatal().LogErrorf("error creating database", err).Err()
	}

	shutdown := func() {
		logger.Info().Log("shutting down the db")
		cancelFunc()
		if err := db.Close(); err != nil {
			logger.Fatal().LogErrorf("error closing DB", err)
		}
	}

	backupFiles, _ := ioutil.ReadDir(filepath.Join("migrations"))
	if len(backupFiles) > 0 {
		if err := database.RunMigrations(logger, config); err != nil {
			return nil, shutdown, logger.Fatal().LogErrorf("error running migrations", err).Err()
		}
	} else {
		logger.Info().Log("there is no backup files of database")
	}

	logger.Info().Log("finished initializing db")

	return db, shutdown, err
}
