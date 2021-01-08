// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package logout

import (
	"bytes"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"testing"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/mockutil"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/token"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
)

func TestNewLogoutCommand(t *testing.T) {
	type args struct {
		cfg        *config.Config
		connection *connection.Connection
	}
	tests := []struct {
		name             string
		args             args
		wantAccessToken  string
		wantRefreshToken string
	}{
		{
			name:             "Successfully logs out",
			wantAccessToken:  "",
			wantRefreshToken: "",
			args: args{
				cfg: &config.Config{
					AccessToken:  "valid",
					RefreshToken: "valid",
				},
				connection: &connection.Connection{
					Token: &token.Token{
						AccessToken:  "valid",
						RefreshToken: "valid",
					},
				},
			},
		},
		{
			name:             "Log out is unsuccessful when tokens are expired",
			wantAccessToken:  "expired",
			wantRefreshToken: "expired",
			args: args{
				cfg: &config.Config{
					AccessToken:  "expired",
					RefreshToken: "expired",
				},
				connection: &connection.Connection{
					Token: &token.Token{
						AccessToken:  "expired",
						RefreshToken: "expired",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		// nolint
		t.Run(tt.name, func(t *testing.T) {
			factory := &factory.Factory{
				Config: mockutil.NewConfigMock(tt.args.cfg),
				Connection: func() (connection.IConnection, error) {
					return mockutil.NewConnectionMock(tt.args.connection, nil), nil
				},
				Logger: func() (logging.Logger, error) {
					loggerBuilder := logging.NewStdLoggerBuilder()
					loggerBuilder = loggerBuilder.Debug(true)
					logger, err := loggerBuilder.Build()
					if err != nil {
						return nil, err
					}

					return logger, nil
				},
			}

			cmd := NewLogoutCommand(factory)
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			_ = cmd.Execute()

			cfg, _ := factory.Config.Load()
			if cfg.AccessToken != tt.wantAccessToken && cfg.RefreshToken != tt.wantRefreshToken {
				t.Errorf("Expected access token and refresh tokens to be cleared in config")
			}
		})
	}
}
