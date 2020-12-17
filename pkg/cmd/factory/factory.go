package factory

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
)

type Factory struct {
	Config     config.IConfig
	Connection func() (connection.IConnection, error)
}
