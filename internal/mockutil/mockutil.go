package mockutil

import (
	"context"
	"errors"

	kafkamgmtv1 "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/api"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
)

func NewConfigMock(cfg *config.Config) config.IConfig {
	return &config.IConfigMock{
		LocationFunc: func() (string, error) {
			return ":mock_location:", nil
		},
		LoadFunc: func() (*config.Config, error) {
			return cfg, nil
		},
		SaveFunc: func(c *config.Config) error {
			cfg = c
			return nil
		},
		RemoveFunc: func() error {
			cfg = nil
			return nil
		},
	}
}

func NewConnectionMock(conn *connection.KeycloakConnection, apiClient *kafkamgmtv1.APIClient) connection.Connection {
	return &connection.ConnectionMock{
		RefreshTokensFunc: func(ctx context.Context) error {
			if conn.Token.AccessToken == "" && conn.Token.RefreshToken == "" {
				return errors.New("")
			}
			if conn.Token.RefreshToken == "expired" {
				return errors.New("")
			}

			return nil
		},
		LogoutFunc: func(ctx context.Context) error {
			if conn.Token.AccessToken == "" && conn.Token.RefreshToken == "" {
				return errors.New("")
			}
			if conn.Token.AccessToken == "expired" && conn.Token.RefreshToken == "expired" {
				return errors.New("")
			}

			cfg, err := conn.Config.Load()
			if err != nil {
				return err
			}

			cfg.AccessToken = ""
			cfg.RefreshToken = ""
			cfg.MasAccessToken = ""
			cfg.MasRefreshToken = ""

			if err = conn.Config.Save(cfg); err != nil {
				return err
			}

			return nil
		},
		APIFunc: func() *api.API {
			a := &api.API{
				Kafka: func() kafkamgmtv1.DefaultApi {
					return apiClient.DefaultApi
				},
			}

			return a
		},
	}
}

func NewKafkaRequestTypeMock(name string) kafkamgmtv1.KafkaRequest {
	var kafkaReq kafkamgmtv1.KafkaRequest
	kafkaReq.SetId("1")
	kafkaReq.SetName(name)

	return kafkaReq
}
