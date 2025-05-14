package config

import (
	"log"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App AppConfig `yaml:"app"`
		GRPC GRPCConfig `yaml:"grpc"`
		Log LogConfig `yaml:"logger"`
		Auth AuthConfig `yaml:"auth"`
		User UserConfig `yaml:"user"`
		Gpt GptConfig `yaml:"gpt"`
	}

	AppConfig struct {
		Name string `yaml:"name"`
		Version string `yaml:"version"`
	}

	GRPCConfig struct {
		Port string `yaml:"port"`
		Timeout int `yaml:"timeout"`
	}

	LogConfig struct {
		Level string `yaml:"level"`
	}

	AuthConfig struct {
		Port string `yaml:"port"`
		Timeout int `yaml:"timeout"`
	}

	UserConfig struct {
		Port string `yaml:"port"`
		Timeout int `yaml:"timeout"`
	}

	GptConfig struct {
		Port string `yaml:"port"`
		Timeout int `yaml:"timeout"`
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