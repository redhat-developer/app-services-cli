package logout

import (
	"bytes"
	"testing"

	"github.com/redhat-developer/app-services-cli/pkg/localize/goi18n"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/redhat-developer/app-services-cli/pkg/connection"

	"github.com/redhat-developer/app-services-cli/internal/mockutil"

	"github.com/redhat-developer/app-services-cli/internal/config"

	"github.com/redhat-developer/app-services-cli/pkg/auth/token"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
)

func TestNewLogoutCommand(t *testing.T) {
	localizer, _ := goi18n.New(nil)
	type args struct {
		cfg        *config.Config
		connection *connection.KeycloakConnection
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
				connection: &connection.KeycloakConnection{
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
				connection: &connection.KeycloakConnection{
					Token: &token.Token{
						AccessToken:  "expired",
						RefreshToken: "expired",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.args.connection.Config = mockutil.NewConfigMock(tt.args.cfg)

		loggerBuilder := logging.NewStdLoggerBuilder()
		loggerBuilder = loggerBuilder.Debug(true)
		logger, _ := loggerBuilder.Build()

		t.Run(tt.name, func(t *testing.T) {
			fact := &factory.Factory{
				Config: mockutil.NewConfigMock(tt.args.cfg),
				Connection: func(connectionCfg *connection.Config) (connection.Connection, error) {
					return mockutil.NewConnectionMock(tt.args.connection, nil), nil
				},
				Localizer: localizer,
				Logger:    logger,
			}

			cmd := NewLogoutCommand(fact)
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			_ = cmd.Execute()

			cfg, _ := fact.Config.Load()
			if cfg.AccessToken != tt.wantAccessToken && cfg.RefreshToken != tt.wantRefreshToken {
				t.Errorf("Expected access token and refresh tokens to be cleared in config")
			}
		})
	}
}