package config

import (
	"path/filepath"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"encoding/json"
	"os"
	"fmt"
)

// Config is the type used to track the config of the client
type Config struct {
	User     string           `json:"user"`
	Services ServiceConfigMap `json:"services,omitempty"`
}

// ServiceConfigMap is a map of configs for the managed application services
type ServiceConfigMap struct {
	Kafka KafkaConfig `json:"kafka,omitempty"`
}

// KafkaConfig is the config for the managed Kafka service
type KafkaConfig struct {
	ClusterID string `json:"clusterId"`
}

// SetKafka sets the current Kafka cluster
func (s *ServiceConfigMap) SetKafka(k *KafkaConfig) {
	s.Kafka = *k
}

// Save saves the given configuration to the configuration file.
func Save(cfg *Config) error {
	file, err := Location()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("can't marshal config: %v", err)
	}
	err = ioutil.WriteFile(file, data, 0600)
	if err != nil {
		return fmt.Errorf("can't write file '%s': %v", file, err)
	}
	return nil
}

// Load loads the configuration from the configuration file. If the configuration file doesn't exist
// it will return an empty configuration object.
func Load() (cfg *Config, err error) {
	file, err := Location()
	if err != nil {
		return
	}
	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		cfg = nil
		err = nil
		return
	}
	if err != nil {
		err = fmt.Errorf("can't check if config file '%s' exists: %v", file, err)
		return
	}
	// #nosec G304
	data, err := ioutil.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("can't read config file '%s': %v", file, err)
		return
	}
	cfg = new(Config)
	err = json.Unmarshal(data, cfg)
	if err != nil {
		err = fmt.Errorf("can't parse config file '%s': %v", file, err)
		return
	}
	return
}

// Location returns the location of the configuration file.
func Location() (path string, err error) {
	if rhmasConfig := os.Getenv("RHMASCLI_CONFIG"); rhmasConfig != "" {
		path = rhmasConfig
	} else {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(home, ".rhmascli.json")
	}
	return path, nil
}