package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"golang.org/x/time/rate"
)

const (
	// DefaultConfigPath is the default path to the config file
	DefaultConfigPath = "/etc/uptimo.toml"
)

type DBConfig struct {
	Driver string `toml:"driver"`
	URL    string `toml:"url"`
}

type ServerConfig struct {
	Port      int        `toml:"port"`
	RateLimit rate.Limit `toml:"rate_limit"`
}

type Config struct {
	DB     DBConfig     `toml:"db"`
	Debug  bool         `toml:"debug"`
	Server ServerConfig `toml:"server"`
}

func readConfig(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}

	return &config, nil 
}

func LoadConfig() (*Config, error) {
	if path, ok := os.LookupEnv("UPTIMO_CONFIG_PATH"); ok {
		return readConfig(path)
	} else {
		return readConfig(DefaultConfigPath)
	}
}
