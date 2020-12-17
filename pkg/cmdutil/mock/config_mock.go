package mock

import "github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"

type Config struct {
	Cfg *config.Config
}

func (c *Config) Load() (*config.Config, error) {
	return c.Cfg, nil
}

func (c *Config) Save(config *config.Config) error {
	c.Cfg = config
	return nil
}

func (c *Config) Remove() error {
	c.Cfg = nil
	return nil
}

func (c *Config) Location() (string, error) {
	return ":inmemory:", nil
}
