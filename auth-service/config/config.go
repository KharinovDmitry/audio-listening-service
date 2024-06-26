package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type HttpConfig struct {
	Port int `yaml:"port"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type Config struct {
	Env            string        `yaml:"env"`
	DBConnStr      string        `yaml:"db_conn_str"`
	MBConnStr      string        `yaml:"mb_conn_str"`
	JwtSecret      string        `yaml:"jwt_secret"`
	Salt           string        `yaml:"salt"`
	MigrationsPath string        `yaml:"migrations_path"`
	TokenTTL       time.Duration `yaml:"token_ttl"`
	TimeoutDB      time.Duration `yaml:"timeout_db"`
	Http           HttpConfig    `yaml:"http"`
	GRPC           GRPCConfig    `yaml:"grpc"`
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
