package create

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"net/http"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/mockutil"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/token"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"gopkg.in/yaml.v2"
)

// nolint:funlen
func TestNewCreateCommand(t *testing.T) {
	type args struct {
		f            *factory.Factory
		name         string
		outputFormat string
		isTTY        bool
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantFormat string
	}{
		{
			name:    "Throw error when stdin is closed and no --name flag value is passed",
			wantErr: true,
			args: args{
				name:         "",
				isTTY:        false,
				outputFormat: "json",
				f: &factory.Factory{
					Logger: func() (logging.Logger, error) {
						loggerBuilder := logging.NewStdLoggerBuilder()
						loggerBuilder = loggerBuilder.Debug(true)
						logger, err := loggerBuilder.Build()
						if err != nil {
							return nil, err
						}

						return logger, nil
					},
					Config: mockutil.NewConfigMock(&config.Config{
						AccessToken:  "valid",
						RefreshToken: "valid",
					}),
					Connection: func() (connection.Connection, error) {
						mockDefaultAPI := &managedservices.DefaultApiMock{}
						mockDefaultAPI.CreateKafkaFunc = func(ctx context.Context) managedservices.ApiCreateKafkaRequest {
							req := managedservices.ApiCreateKafkaRequest{
								ApiService: mockDefaultAPI,
							}

							return req
						}
						mockDefaultAPI.CreateKafkaExecuteFunc = func(r managedservices.ApiCreateKafkaRequest) (managedservices.KafkaRequest, *http.Response, managedservices.GenericOpenAPIError) {
							kafkaReq := mockutil.NewKafkaRequestTypeMock("test-kafka")
							var genericError managedservices.GenericOpenAPIError
							return kafkaReq, nil, genericError
						}
						apiClient := &managedservices.APIClient{
							DefaultApi: mockDefaultAPI,
						}

						mockConnection := mockutil.NewConnectionMock(&connection.KeycloakConnection{
							Token: &token.Token{
								AccessToken:  "valid",
								RefreshToken: "valid",
							},
						}, apiClient)

						return mockConnection, nil
					},
				},
			},
		},
		{
			name: "Create Kafka request and output as JSON",
			args: args{
				name:         "test-kafka",
				outputFormat: "json",
				isTTY:        true,
				f: &factory.Factory{
					Logger: func() (logging.Logger, error) {
						loggerBuilder := logging.NewStdLoggerBuilder()
						loggerBuilder = loggerBuilder.Debug(true)
						logger, err := loggerBuilder.Build()
						if err != nil {
							return nil, err
						}

						return logger, nil
					},
					Config: mockutil.NewConfigMock(&config.Config{
						AccessToken:  "valid",
						RefreshToken: "valid",
					}),
					Connection: func() (connection.Connection, error) {
						mockDefaultAPI := &managedservices.DefaultApiMock{}
						mockDefaultAPI.CreateKafkaFunc = func(ctx context.Context) managedservices.ApiCreateKafkaRequest {
							req := managedservices.ApiCreateKafkaRequest{
								ApiService: mockDefaultAPI,
							}

							return req
						}
						mockDefaultAPI.CreateKafkaExecuteFunc = func(r managedservices.ApiCreateKafkaRequest) (managedservices.KafkaRequest, *http.Response, managedservices.GenericOpenAPIError) {
							kafkaReq := mockutil.NewKafkaRequestTypeMock("test-kafka")
							var genericError managedservices.GenericOpenAPIError
							return kafkaReq, nil, genericError
						}
						apiClient := &managedservices.APIClient{
							DefaultApi: mockDefaultAPI,
						}

						mockConnection := mockutil.NewConnectionMock(&connection.KeycloakConnection{
							Token: &token.Token{
								AccessToken:  "valid",
								RefreshToken: "valid",
							},
						}, apiClient)

						return mockConnection, nil
					},
				},
			},
		},
		{
			name: "Create Kafka request and output as YAML",
			args: args{
				isTTY:        true,
				name:         "test",
				outputFormat: "yaml",
				f: &factory.Factory{
					Logger: func() (logging.Logger, error) {
						loggerBuilder := logging.NewStdLoggerBuilder()
						loggerBuilder = loggerBuilder.Debug(true)
						logger, err := loggerBuilder.Build()
						if err != nil {
							return nil, err
						}

						return logger, nil
					},
					Config: mockutil.NewConfigMock(&config.Config{
						AccessToken:  "valid",
						RefreshToken: "valid",
					}),
					Connection: func() (connection.Connection, error) {
						mockDefaultAPI := &managedservices.DefaultApiMock{}
						mockDefaultAPI.CreateKafkaFunc = func(ctx context.Context) managedservices.ApiCreateKafkaRequest {
							req := managedservices.ApiCreateKafkaRequest{
								ApiService: mockDefaultAPI,
							}

							return req
						}
						mockDefaultAPI.CreateKafkaExecuteFunc = func(r managedservices.ApiCreateKafkaRequest) (managedservices.KafkaRequest, *http.Response, managedservices.GenericOpenAPIError) {
							kafkaReq := mockutil.NewKafkaRequestTypeMock("test-kafka")
							var genericError managedservices.GenericOpenAPIError
							return kafkaReq, nil, genericError
						}
						apiClient := &managedservices.APIClient{
							DefaultApi: mockDefaultAPI,
						}

						mockConnection := mockutil.NewConnectionMock(&connection.KeycloakConnection{
							Token: &token.Token{
								AccessToken:  "valid",
								RefreshToken: "valid",
							},
						}, apiClient)

						return mockConnection, nil
					},
				},
			},
		},
		{
			name:    "Throw an error when invalid output format is passed",
			wantErr: true,
			args: args{
				isTTY:        true,
				name:         "test-kafka",
				outputFormat: "xml",
				f: &factory.Factory{
					Logger: func() (logging.Logger, error) {
						loggerBuilder := logging.NewStdLoggerBuilder()
						loggerBuilder = loggerBuilder.Debug(true)
						logger, err := loggerBuilder.Build()
						if err != nil {
							return nil, err
						}

						return logger, nil
					},
					Config: mockutil.NewConfigMock(&config.Config{
						AccessToken:  "valid",
						RefreshToken: "valid",
					}),
					Connection: func() (connection.Connection, error) {
						mockDefaultAPI := &managedservices.DefaultApiMock{}
						mockDefaultAPI.CreateKafkaFunc = func(ctx context.Context) managedservices.ApiCreateKafkaRequest {
							req := managedservices.ApiCreateKafkaRequest{
								ApiService: mockDefaultAPI,
							}

							return req
						}
						mockDefaultAPI.CreateKafkaExecuteFunc = func(r managedservices.ApiCreateKafkaRequest) (managedservices.KafkaRequest, *http.Response, managedservices.GenericOpenAPIError) {
							kafkaReq := mockutil.NewKafkaRequestTypeMock("")
							var genericError managedservices.GenericOpenAPIError
							return kafkaReq, nil, genericError
						}
						apiClient := &managedservices.APIClient{
							DefaultApi: mockDefaultAPI,
						}

						mockConnection := mockutil.NewConnectionMock(&connection.KeycloakConnection{
							Token: &token.Token{
								AccessToken:  "valid",
								RefreshToken: "valid",
							},
						}, apiClient)

						return mockConnection, nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		// nolint
		t.Run(tt.name, func(t *testing.T) {

			b := bytes.NewBufferString("")
			tt.args.f.IOStreams = &iostreams.IOStreams{
				Out: b,
			}

			cmd := NewCreateCommand(tt.args.f)
			cmd.SetArgs([]string{
				"--name",
				tt.args.name,
				"--output",
				tt.args.outputFormat,
			})

			err := cmd.Execute()
			if !tt.wantErr && err != nil {
				t.Fatal("Expected error but got nil")
				return
			}

			if tt.args.name == "" && !tt.args.isTTY && (tt.wantErr == (err == nil)) {
				t.Fatalf("wantErr = %v, err = %v", tt.wantErr, err)
				return
			} else if !tt.args.isTTY && tt.args.name == "" {
				return
			}

			out, _ := ioutil.ReadAll(b)

			var kafkaRequest managedservices.KafkaRequest
			switch tt.args.outputFormat {
			case "json":
				err := json.Unmarshal(out, &kafkaRequest)
				if err != nil {
					t.Fatalf("Expected JSON output: %v", err.Error())
					return
				}
			case "yaml", "yml":
				err := yaml.Unmarshal(out, &kafkaRequest)
				if err != nil {
					t.Fatalf("Expected YAML output: %v", err.Error())
					return
				}
			}
		})
	}
}
