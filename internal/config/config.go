package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const (
	configPath = "././configs/config.yaml"
)

type Config struct {
	ConnectionString string `yaml:"connectionString"`
	Address          string `yaml:"address"`
	JwtSecret        string `yaml:"address"`
	TokenTimeToLive  string `yaml:"address"`
}

func LoadConfig() *Config {
	config := &Config{}
	file, err := os.ReadFile(configPath)

	if err != nil {
		panic("Config not found!")
	}

	err = yaml.Unmarshal(file, config)

	if err != nil {
		panic("Could not unmarshal config correct.")
	}

	return config
}
