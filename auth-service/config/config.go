package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Env            string        `yaml:"env"`
	ConnStr        string        `yaml:"conn_str"`
	GRPC           GRPCConfig    `yaml:"grpc"`
	MigrationsPath string        `yaml:"migrations_path"`
	TokenTTL       time.Duration `yaml:"token_ttl"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func Load(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	cfgBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err = yaml.Unmarshal(cfgBytes, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
