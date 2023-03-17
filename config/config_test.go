package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	if _, err := os.Stat(DefaultConfigPath); os.IsNotExist(err) {
		t.Errorf("Default config file not found")
	}

	config, err := LoadConfig()

	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}

	assert.Equal(t, 8080, config.Server.Port)
}

func TestLoadConfigOverride(t *testing.T) {
	t.Run("prepare", func(t *testing.T) {
		if err := os.Setenv("UPTIMO_CONFIG_PATH", "test/data/config.toml"); err != nil {
			t.Errorf("Error setting env variable: %v", err)
		}
	})

	config, err := LoadConfig()

	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}

	// Different from the default config file
	assert.Equal(t, 4343, config.Server.Port)
}
