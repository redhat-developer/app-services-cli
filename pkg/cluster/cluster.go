package cluster

import (
	"context"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
)

type CommandConnectOptions struct {
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
	IO         *iostreams.IOStreams

	KubeconfigLocation string
	OfflineAccessToken string
	Namespace          string
	Force              bool
	IgnoreConfig       bool

	SelectedKafka string
}

// Cluster defines methods used to interact with a cluster
type Cluster interface {
	Connect(ctx context.Context, opts *CommandConnectOptions) error
	IsKafkaConnectionCRDInstalled(ctx context.Context) (bool, error)
	CurrentNamespace() (string, error)
}
