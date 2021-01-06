package create

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

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
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantFormat string
	}{
		{
			name: "Create Kafka request and output as JSON",
			args: args{
				name:         "test-kafka",
				outputFormat: "json",
				f: &factory.Factory{
					Config: mockutil.NewConfigMock(&config.Config{
						AccessToken:  "valid",
						RefreshToken: "valid",
					}),
					Connection: func() (connection.IConnection, error) {
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

						mockConnection := mockutil.NewConnectionMock(&connection.Connection{
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
				name:         "test-kafka",
				outputFormat: "yaml",
				f: &factory.Factory{
					Config: mockutil.NewConfigMock(&config.Config{
						AccessToken:  "valid",
						RefreshToken: "valid",
					}),
					Connection: func() (connection.IConnection, error) {
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

						mockConnection := mockutil.NewConnectionMock(&connection.Connection{
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
				name:         "test-kafka",
				outputFormat: "xml",
				f: &factory.Factory{
					Config: mockutil.NewConfigMock(&config.Config{
						AccessToken:  "valid",
						RefreshToken: "valid",
					}),
					Connection: func() (connection.IConnection, error) {
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

						mockConnection := mockutil.NewConnectionMock(&connection.Connection{
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
			cmd := NewCreateCommand(tt.args.f)
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
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
