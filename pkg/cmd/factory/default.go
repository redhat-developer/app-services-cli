package factory

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"context"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/debug"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
)

// New creates a new command factory
// The command factory is available to all command packages
// giving centralized access to the config and API connection
// nolint:funlen
func New(cliVersion string) *Factory {
	io := iostreams.System()

	var logger logging.Logger
	var conn connection.Connection
	cfgFile := config.NewFile()

	loggerFunc := func() (logging.Logger, error) {
		if logger != nil {
			return logger, nil
		}

		loggerBuilder := logging.NewStdLoggerBuilder()
		loggerBuilder = loggerBuilder.Streams(io.Out, io.ErrOut)

		debugEnabled := debug.Enabled()
		loggerBuilder = loggerBuilder.Debug(debugEnabled)

		logger, err := loggerBuilder.Build()

		if err != nil {
			return nil, err
		}

		return logger, nil
	}

	connectionFunc := func() (connection.Connection, error) {
		if conn != nil {
			return conn, nil
		}

		cfg, err := cfgFile.Load()
		if err != nil {
			return nil, err
		}

		builder := connection.NewBuilder()

		if cfg.AccessToken != "" {
			builder.WithAccessToken(cfg.AccessToken)
		}
		if cfg.RefreshToken != "" {
			builder.WithRefreshToken(cfg.RefreshToken)
		}
		if cfg.ClientID != "" {
			builder.WithClientID(cfg.ClientID)
		}
		if cfg.Scopes != nil {
			builder.WithScopes(cfg.Scopes...)
		}
		if cfg.APIGateway != "" {
			builder.WithURL(cfg.APIGateway)
		}
		if cfg.AuthURL == "" {
			cfg.AuthURL = connection.DefaultAuthURL
		}
		builder.WithAuthURL(cfg.AuthURL)

		builder.WithInsecure(cfg.Insecure)

		// create a logger if it has not already been created
		logger, err = loggerFunc()
		if err != nil {
			return nil, err
		}

		conn, err = builder.Build()
		if err != nil {
			return nil, err
		}

		accessTk, refreshTk, err := conn.RefreshTokens(context.TODO())
		if err != nil {
			return nil, err
		}

		accessTkChanged := accessTk != cfg.AccessToken
		refreshTkChanged := refreshTk != cfg.RefreshToken

		if accessTkChanged {
			cfg.AccessToken = accessTk
		}
		if refreshTkChanged {
			cfg.RefreshToken = refreshTk
		}

		if !accessTkChanged && refreshTkChanged {
			return conn, nil
		}

		err = cfgFile.Save(cfg)
		if err != nil {
			logger.Debug(localizer.MustLocalizeFromID("common.log.debug.couldNotSaveRefreshTokenToConfig"), err)
		}

		return conn, nil
	}

	return &Factory{
		IOStreams:  io,
		Config:     cfgFile,
		Connection: connectionFunc,
		Logger:     loggerFunc,
	}
}
