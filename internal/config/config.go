package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Database `yaml:"database"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func NewConfig() *Config {
	return &Config{
		Database{
			Host:     "",
			Port:     "",
			Password: "",
			User:     "",
		},
	}
}

func LoadConfig() *Config {
	config := NewConfig()
	file, err := os.ReadFile("././config/app.yaml")

	if err != nil {
		panic("Config not found!")
	}

	err = yaml.Unmarshal(file, config)

	if err != nil {
		log.Fatalf("Config con not be unmarshal: %v", err)
	}

	return config
}
