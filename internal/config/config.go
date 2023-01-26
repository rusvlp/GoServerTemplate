package config

import (
	"flag"
	"github.com/BurntSushi/toml"
)

var configPath string

type Config struct {
	Port string `toml:"port"`
}

func new() *Config {
	return &Config{
		Port: "8081",
	}
}

func Initialize() (*Config, error) {
	flag.StringVar(&configPath, "cfg", "config/config.toml", "path to file of configuration")
	flag.Parse()
	config := new()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		return nil, err
	}

	cfgerr := config.checkConfigErrors()

	if cfgerr != nil {
		return nil, cfgerr
	}

	return config, nil
}

func (config *Config) checkConfigErrors() error {
	return nil // Допилить чек конфига на ошибки
}
