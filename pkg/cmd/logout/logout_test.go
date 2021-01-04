// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package logout

import (
	"bytes"
	"testing"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/mock"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
)

func TestNewLogoutCommand(t *testing.T) {
	type args struct {
		f *factory.Factory
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
				f: &factory.Factory{
					Config: &mock.Config{
						Cfg: &config.Config{
							AccessToken:  "valid",
							RefreshToken: "valid",
						},
					},
					Connection: func() (connection.IConnection, error) {
						return &mock.Connection{
							AccessToken:  "valid",
							RefreshToken: "valid",
						}, nil
					},
				},
			},
		},
		{
			name:             "Log out fails when tokens are expired",
			wantAccessToken:  "expired",
			wantRefreshToken: "expired",
			args: args{
				f: &factory.Factory{
					Config: &mock.Config{
						Cfg: &config.Config{
							AccessToken:  "expired",
							RefreshToken: "expired",
						},
					},
					Connection: func() (connection.IConnection, error) {
						return &mock.Connection{
							AccessToken:  "expired",
							RefreshToken: "expired",
						}, nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		// nolint
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewLogoutCommand(tt.args.f)
			b := bytes.NewBufferString("")
			cmd.SetOut(b)

			_ = cmd.Execute()

			cfg, _ := tt.args.f.Config.Load()
			if cfg.AccessToken != tt.wantAccessToken && cfg.RefreshToken != tt.wantRefreshToken {
				t.Errorf("Expected access token and refresh tokens to be cleared in config")
			}
		})
	}
}
