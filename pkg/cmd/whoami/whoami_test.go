package whoami

import (
	"bytes"
	"strings"
	"testing"

	"github.com/redhat-developer/app-services-cli/internal/mockutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/auth/token"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize/goi18n"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/kcconnection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
)

func TestNewWhoAmICmd(t *testing.T) {

	type args struct {
		cfg        *config.Config
		connection *kcconnection.Connection
	}

	tests := []struct {
		name             string
		args             args
		expectedUserName string
		wantErr          bool
	}{
		{
			name: "Successfully shows username",
			args: args{
				cfg: &config.Config{
					AccessToken:  JDoeAccessToken,
					RefreshToken: "valid",
				},
				connection: &kcconnection.Connection{
					Token: &token.Token{
						AccessToken:  "valid",
						RefreshToken: "valid",
					},
				},
			},
			expectedUserName: "jdoe",
			wantErr:          false,
		},
		{
			name: "Successfully fails while displaying username",
			args: args{
				cfg: &config.Config{
					AccessToken:  "..LcHxtPnO482VKC0H_1x",
					RefreshToken: "valid",
				},
				connection: &kcconnection.Connection{
					Token: &token.Token{
						AccessToken:  "valid",
						RefreshToken: "valid",
					},
				},
			},
			expectedUserName: "valid",
			wantErr:          true,
		},
		{
			name: "Should print empty string if invalid token is passed",
			args: args{
				cfg: &config.Config{
					AccessToken:  "..LcHxtPnO482VKC0H_1x",
					RefreshToken: "valid",
				},
				connection: &kcconnection.Connection{
					Token: &token.Token{
						AccessToken:  "valid",
						RefreshToken: "valid",
					},
				},
			},
			expectedUserName: "",
			wantErr:          false,
		},
		{
			name: "Should fail if wrong name is displayed",
			args: args{
				cfg: &config.Config{
					AccessToken:  HWorldAccessToken,
					RefreshToken: "valid",
				},
				connection: &kcconnection.Connection{
					Token: &token.Token{
						AccessToken:  "valid",
						RefreshToken: "valid",
					},
				},
			},
			expectedUserName: "hellow",
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		tt.args.connection.Config = mockutil.NewConfigMock(tt.args.cfg)

		loggerBuilder := logging.NewStdLoggerBuilder()
		loggerBuilder = loggerBuilder.Debug(true)
		logger, _ := loggerBuilder.Build()
		localizer, _ := goi18n.New(nil)

		t.Run(tt.name, func(t *testing.T) {

			writerStream := iostreams.System()

			buf := bytes.NewBufferString("")
			writerStream.Out = buf

			fact := &factory.Factory{
				IOStreams: writerStream,
				Config:    mockutil.NewConfigMock(tt.args.cfg),
				Connection: func() (connection.Connection, error) {
					return mockutil.NewConnectionMock(tt.args.connection, nil), nil
				},
				Localizer: localizer,
				Logger:    logger,
			}

			cmd := NewWhoAmICmd(fact)

			cmd.SetOut(buf)

			_ = cmd.Execute()

			bufStr := strings.Trim(buf.String(), "\n")

			if (bufStr != tt.expectedUserName) != tt.wantErr {
				t.Errorf("Expected username = \"%s\", fetched username = \"%s\"", tt.expectedUserName, bufStr)
			}
		})
	}
}
