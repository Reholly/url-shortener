package config

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

type Config struct {
	ConnectionString string `yaml:"connectionString"`
	MigrationsPath   string `yaml:"migrationsPath"`
	Address          string `yaml:"address"`
}

func LoadConfig() *Config {
	config := &Config{}
	file, err := os.ReadFile("././config/app.yaml")

	if err != nil {
		panic("Config not found!")
	}

	err = yaml.Unmarshal(file, config)

	if err != nil {
		slog.Error("Config con not be unmarshal: %v", err)
	}

	return config
}
