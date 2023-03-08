package delete

import (
	"bytes"
	"fmt"

	"testing"

	"github.com/redhat-developer/app-services-cli/internal/mockutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/auth/token"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize/goi18n"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/kcconnection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
)

func TestNewDeleteCommand(t *testing.T) {
	localizer, _ := goi18n.New(nil)
	type args struct {
		context *servicecontext.Context
	}
	tests := []struct {
		name          string
		args          args
		targetContext string
		wantErr       bool
	}{
		{
			name: "Successfully deletes an existing context",
			args: args{
				context: &servicecontext.Context{
					Contexts: map[string]servicecontext.ServiceConfig{
						"c1": {
							KafkaID: "my-kafka-id-1",
						},
						"c2": {
							KafkaID: "my-kafka-id-2",
						},
					},
				},
			},
			targetContext: "c2",
			wantErr:       false,
		},
		{
			name: "Should fail if a non-existing context is being deleted",
			args: args{
				context: &servicecontext.Context{
					Contexts: map[string]servicecontext.ServiceConfig{
						"c1": {
							KafkaID: "my-kafka-id-1",
						},
						"c2": {
							KafkaID: "my-kafka-id-2",
						},
					},
				},
			},
			targetContext: "c3",
			wantErr:       true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			loggerBuilder := logging.NewStdLoggerBuilder()
			loggerBuilder = loggerBuilder.Debug(true)
			logger, _ := loggerBuilder.Build()

			writerStream := iostreams.System()

			buf := bytes.NewBufferString("")
			writerStream.Out = buf

			conn := &kcconnection.Connection{
				Token: &token.Token{
					AccessToken:  "valid",
					RefreshToken: "valid",
				},
			}

			fact := &factory.Factory{
				IOStreams:      writerStream,
				ServiceContext: mockutil.NewContextMock(tt.args.context),
				Connection: func() (connection.Connection, error) {
					return mockutil.NewConnectionMock(conn, nil), nil
				},
				Localizer: localizer,
				Logger:    logger,
			}

			cmd := NewDeleteCommand(fact)

			nameFlag := fmt.Sprintf("--name=%s", tt.targetContext)
			cmd.SetArgs([]string{nameFlag})
			err := cmd.Execute()

			if (err != nil) != tt.wantErr {
				t.Errorf("Deleting context \"%s\" throws error: %v", tt.targetContext, err)
			}
		})
	}
}
