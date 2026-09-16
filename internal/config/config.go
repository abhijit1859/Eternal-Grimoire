package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type StorageConfig struct {
	Type      string `yaml:"type"`
	Enabled   bool   `yaml:"enabled"`
	Path      string `yaml:"path"`
	Bucket    string `yaml:"bucket"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Region    string `yaml:"region"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Storage  StorageConfig  `yaml:"storage"`
}

func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return Config{}, fmt.Errorf("failed to read the config file")
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("Failed to parse yaml file")
	}

	return cfg, nil
}

 