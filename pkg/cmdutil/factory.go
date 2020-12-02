package cmdutil

import "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"

type Factory struct {
	Config func() (config.Config, error)
}
