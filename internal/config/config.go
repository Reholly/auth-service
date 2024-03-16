package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const (
	configPath = "././configs/config.yaml"
)

type Config struct {
	ConnectionString       string `yaml:"connectionString"`
	Address                string `yaml:"address"`
	JwtSecret              string `yaml:"jwt_secret"`
	TokenTimeToLiveInHours int    `yaml:"token_time_to_live_in_hours"`
	Salt                   string `yaml:"salt"`

	SmtpFrom                 string `yaml:"smtp_from"`
	SmtpPassword             string `yaml:"smtp_password"`
	SmtpHost                 string `yaml:"smtp_host"`
	SmtpPort                 string `yaml:"smtp_port"`
	EmailConfirmationUrlBase string `yaml:"email_confirmation_url_base"`
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
