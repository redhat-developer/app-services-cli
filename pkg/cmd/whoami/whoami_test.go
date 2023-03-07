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
					AccessToken:  "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ii00ZWxjX1ZkTl9Xc09VWWYyRzRReHI4R2N3SXhfS3RYVUNpdGF0TEtsTHcifQ.eyJleHAiOjE2NzAzMTE3MDQsImlhdCI6MTY3MDMxMDgwNCwiYXV0aF90aW1lIjoxNjcwMzA0NDE4LCJqdGkiOiI4YTI2ZTI3Yy1iYzdmLTRiNWEtYWI3ZC1mN2I5MmEwYWZkYTgiLCJpc3MiOiJodHRwczovL3Nzby5yZWRoYXQuY29tL2F1dGgvcmVhbG1zL3JlZGhhdC1leHRlcm5hbCIsImF1ZCI6InJob2FzLWNsaS1wcm9kIiwic3ViIjoiZjo1MjhkNzZmZi1mNzA4LTQzZWQtOGNkNS1mZTE2ZjRmZTBjZTY6cnBhdHRuYWkiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJyaG9hcy1jbGktcHJvZCIsInNlc3Npb25fc3RhdGUiOiI5YTkwNmI5My0zNjUwLTRkODEtYWEzZC0zZDQwOWZmMWE5ZTYiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cDovLzEyNy4wLjAuMSIsImh0dHA6Ly9sb2NhbGhvc3QiLCJodHRwczovL2NvbnNvbGUucmVkaGF0LmNvbSJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiYXV0aGVudGljYXRlZCIsInBvcnRhbF9tYW5hZ2Vfc3Vic2NyaXB0aW9ucyIsImVycmF0YTpub3RpZmljYXRpb25fc3RhdHVzX2VuYWJsZWQiLCJlcnJhdGE6bm90aWZpY2F0aW9uOmVuaGFuY2VtZW50IiwicG9ydGFsX21hbmFnZV9jYXNlcyIsImVycmF0YTpub3RpZmljYXRpb246c2VjdXJpdHkiLCJlcnJhdGE6bm90aWZpY2F0aW9uX2xldmVsX3N5c3RlbS12aXNpYmxlIiwiZXJyYXRhOm5vdGlmaWNhdGlvbjpidWdmaXgiLCJvZmZsaW5lX2FjY2VzcyIsImFkbWluOm9yZzphbGwiLCJ1bWFfYXV0aG9yaXphdGlvbiIsInBvcnRhbF9zeXN0ZW1fbWFuYWdlbWVudCIsImVycmF0YTpub3RpZmljYXRpb25fZGVsaXZlcnlfaW5zdGFudCIsInJoZF9hY2Nlc3NfbWlkZGxld2FyZSIsInBvcnRhbF9kb3dubG9hZCJdfSwicmVzb3VyY2VfYWNjZXNzIjp7InJoZC1kbSI6eyJyb2xlcyI6WyJyaHVzZXIiXX0sImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwgYXBpLmlhbS5zZXJ2aWNlX2FjY291bnRzIiwic2lkIjoiOWE5MDZiOTMtMzY1MC00ZDgxLWFhM2QtM2Q0MDlmZjFhOWU2IiwiYWNjb3VudF9udW1iZXIiOiIxMjM0NTY3OCIsImlzX2ludGVybmFsIjpmYWxzZSwiZW1haWxfdmVyaWZpZWQiOnRydWUsInByZWZlcnJlZF91c2VybmFtZSI6Impkb2UiLCJsb2NhbGUiOiJlbl9VUyIsImdpdmVuX25hbWUiOiJKb2huIiwiaXNfb3JnX2FkbWluIjp0cnVlLCJhY2NvdW50X2lkIjoiMTIzNDU2NzgiLCJvcmdfaWQiOiIxMjM0NTY3OCIsInJoLXVzZXItaWQiOiIxMjM0NTY3OCIsInJoLW9yZy1pZCI6IjEyMzQ1Njc4IiwibmFtZSI6IkpvaG4gRG9lIiwiZmFtaWx5X25hbWUiOiJEb2UiLCJlbWFpbCI6Impkb2VAcmFuZG9tLmNvbSJ9.Oh3ZaEEqZ0JmnnN6dGD4-C0zHIo-7XlKyjhb5qg0Cm6UTHTXvbXdSQP8_5nsGioxj8_hY5zIvxCTeb2jDXm5F11JxZJMe66Yd9NX2e5sRyBNiaXHnKUF7jqIcgMezy44sNfi-Qn-MfLBZW2tAo8js6CeiYgxX0sarg1qRzz5nz5RpMX7_T8tDVNXrj34jZaQcwTCCAdTldNDowXbbcDo36lLYxhCePJ6SVoaf3Xp6JLgmOM_5O5iWd3jNM_n3WMlcSta44R8W4QLLhfEcwBAO203m22aH-Z-ynga-Y3kkvwZUcko8D80ZNoaQGHBY_Me5V7zBEzW5RIWnwJGcgyFew",
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
					AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NzAzMTA4MDQsImF1dGhfdGltZSI6MTY3MDMwNDQxOCwianRpIjoiOGEyNmUyN2MtYmM3Zi00YjVhLWFiN2QtZjdiOTJhMGFmZGE4IiwiaXNzIjoiaHR0cHM6Ly9zc28ucmVkaGF0LmNvbS9hdXRoL3JlYWxtcy9yZWRoYXQtZXh0ZXJuYWwiLCJhdWQiOiJyaG9hcy1jbGktcHJvZCIsInN1YiI6ImY6NTI4ZDc2ZmYtZjcwOC00M2VkLThjZDUtZmUxNmY0ZmUwY2U2Omh3b3JsZCIsInR5cCI6IkJlYXJlciIsImF6cCI6InJob2FzLWNsaS1wcm9kIiwic2Vzc2lvbl9zdGF0ZSI6IjlhOTA2YjkzLTM2NTAtNGQ4MS1hYTNkLTNkNDA5ZmYxYTllNiIsImFsbG93ZWQtb3JpZ2lucyI6WyJodHRwOi8vMTI3LjAuMC4xIiwiaHR0cDovL2xvY2FsaG9zdCIsImh0dHBzOi8vY29uc29sZS5yZWRoYXQuY29tIl0sInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJhdXRoZW50aWNhdGVkIiwicG9ydGFsX21hbmFnZV9zdWJzY3JpcHRpb25zIiwiZXJyYXRhOm5vdGlmaWNhdGlvbl9zdGF0dXNfZW5hYmxlZCIsImVycmF0YTpub3RpZmljYXRpb246ZW5oYW5jZW1lbnQiLCJwb3J0YWxfbWFuYWdlX2Nhc2VzIiwiZXJyYXRhOm5vdGlmaWNhdGlvbjpzZWN1cml0eSIsImVycmF0YTpub3RpZmljYXRpb25fbGV2ZWxfc3lzdGVtLXZpc2libGUiLCJlcnJhdGE6bm90aWZpY2F0aW9uOmJ1Z2ZpeCIsIm9mZmxpbmVfYWNjZXNzIiwiYWRtaW46b3JnOmFsbCIsInVtYV9hdXRob3JpemF0aW9uIiwicG9ydGFsX3N5c3RlbV9tYW5hZ2VtZW50IiwiZXJyYXRhOm5vdGlmaWNhdGlvbl9kZWxpdmVyeV9pbnN0YW50IiwicmhkX2FjY2Vzc19taWRkbGV3YXJlIiwicG9ydGFsX2Rvd25sb2FkIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsicmhkLWRtIjp7InJvbGVzIjpbInJodXNlciJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgcHJvZmlsZSBlbWFpbCBhcGkuaWFtLnNlcnZpY2VfYWNjb3VudHMiLCJzaWQiOiI5YTkwNmI5My0zNjUwLTRkODEtYWEzZC0zZDQwOWZmMWE5ZTYiLCJhY2NvdW50X251bWJlciI6IjEyMzQ1Njc4IiwiaXNfaW50ZXJuYWwiOmZhbHNlLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiaHdvcmxkIiwibG9jYWxlIjoiZW5fVVMiLCJnaXZlbl9uYW1lIjoiSm9obiIsImlzX29yZ19hZG1pbiI6dHJ1ZSwiYWNjb3VudF9pZCI6IjEyMzQ1Njc4Iiwib3JnX2lkIjoiMTIzNDU2NzgiLCJyaC11c2VyLWlkIjoiMTIzNDU2NzgiLCJyaC1vcmctaWQiOiIxMjM0NTY3OCIsIm5hbWUiOiJIZWxsbyIsImZhbWlseV9uYW1lIjoiV29ybGQiLCJlbWFpbCI6Imh3b3JsZEByYW5kb20uY29tIn0.LSp4CuYBFe6ovHUT1H992M3oPPJb-1sjG6np-08kd8M",
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
