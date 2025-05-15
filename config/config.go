package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type (
	Config struct {
		App    AppConfig    `yaml:"app"`
		Server ServerConfig `yaml:"server"`
		Log    LogConfig    `yaml:"logger"`
		Auth   AuthConfig   `yaml:"auth"`
		User   UserConfig   `yaml:"user"`
		Gpt    GptConfig    `yaml:"gpt"`
	}

	AppConfig struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	ServerConfig struct {
		Port    string `yaml:"port"`
		Timeout int    `yaml:"timeout"`
	}

	LogConfig struct {
		Level string `yaml:"level"`
	}

	AuthConfig struct {
		Address string `yaml:"address"`
		Timeout int    `yaml:"timeout"`
	}

	UserConfig struct {
		Address string `yaml:"address"`
		Timeout int    `yaml:"timeout"`
	}

	GptConfig struct {
		Address string `yaml:"address"`
		Timeout int    `yaml:"timeout"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig("config.yaml", cfg); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return nil, err
	}
	return cfg, nil
}
