package factory

import (
	"context"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/debug"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
)

// New creates a new command factory
// The command factory is available to all command packages
// giving centralized access to the config and API connection
func New(cliVersion string) *Factory {
	var logger logging.Logger
	cfgFile := config.NewFile()

	connectionFunc := func() (connection.IConnection, error) {
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
		if cfg.URL != "" {
			builder.WithURL(cfg.URL)
		}
		if cfg.AuthURL == "" {
			cfg.AuthURL = connection.DefaultAuthURL
		}
		builder.WithAuthURL(cfg.AuthURL)

		builder.WithInsecure(cfg.Insecure)

		conn, err := builder.Build()
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

		// TODO: Warning log on error
		_ = cfgFile.Save(cfg)

		return conn, nil
	}

	loggerFunc := func() (logging.Logger, error) {
		if logger != nil {
			return logger, nil
		}

		loggerBuilder := logging.NewStdLoggerBuilder()
		debugEnabled := debug.Enabled()
		loggerBuilder = loggerBuilder.Debug(debugEnabled)
		logger, err := loggerBuilder.Build()
		if err != nil {
			return nil, err
		}

		return logger, nil
	}

	return &Factory{
		Config:     cfgFile,
		Connection: connectionFunc,
		Logger:     loggerFunc,
	}
}
