package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port string `yaml:"port"`
}

func NewConfig() *Config {
	return &Config{}
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	c := NewConfig()
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
