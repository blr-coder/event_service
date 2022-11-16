package config

import (
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
)

type Config struct {
	AppPort          string `yaml:"app_port"`
	PostgresConnLink string `yaml:"postgres_conn_link"`
}

func NewConfig(configPath string) (*Config, error) {
	var config Config

	configFile, err := os.OpenFile(configPath, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		return nil, err
	}

	return &config, yaml.NewDecoder(configFile).Decode(&config)
}
