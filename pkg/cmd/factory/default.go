package factory

import (
	"context"
	"net/http"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/debug"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/httputil"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

// New creates a new command factory
// The command factory is available to all command packages
// giving centralized access to the config and API connection

// nolint:funlen
func New(localizer localize.Localizer) *Factory {
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

	connectionFunc := func(connectionCfg *connection.Config) (connection.Connection, error) {
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
		if cfg.MasAccessToken != "" {
			builder.WithMASAccessToken(cfg.MasAccessToken)
		}
		if cfg.MasRefreshToken != "" {
			builder.WithMASRefreshToken(cfg.MasRefreshToken)
		}
		if cfg.ClientID != "" {
			builder.WithClientID(cfg.ClientID)
		}
		if cfg.Scopes != nil {
			builder.WithScopes(cfg.Scopes...)
		}
		if cfg.APIUrl != "" {
			builder.WithURL(cfg.APIUrl)
		}
		if cfg.AuthURL == "" {
			cfg.AuthURL = build.ProductionAuthURL
		}
		builder.WithAuthURL(cfg.AuthURL)

		if cfg.MasAuthURL == "" {
			cfg.MasAuthURL = build.ProductionMasAuthURL
		}
		builder.WithMASAuthURL(cfg.MasAuthURL)

		builder.WithInsecure(cfg.Insecure)

		builder.WithConfig(cfgFile)

		// create a logger if it has not already been created
		logger, err = loggerFunc()
		if err != nil {
			return nil, err
		}

		transportWrapper := func(a http.RoundTripper) http.RoundTripper {
			return &httputil.LoggingRoundTripper{
				Proxied: a,
				Logger:  logger,
			}
		}

		builder.WithTransportWrapper(transportWrapper)

		builder.WithConnectionConfig(connectionCfg)

		conn, err = builder.Build()
		if err != nil {
			return nil, err
		}

		err = conn.RefreshTokens(context.TODO())
		if err != nil {
			return nil, err
		}

		return conn, nil
	}

	return &Factory{
		IOStreams:  io,
		Config:     cfgFile,
		Connection: connectionFunc,
		Logger:     loggerFunc,
		Localizer:  localizer,
	}
}
