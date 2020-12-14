package factory

import "github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"

type Factory struct {
	Config func() (config.Config, error)
}