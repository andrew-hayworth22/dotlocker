// Package config provides access to the CLI's configuration
package config

import (
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	RepoURL string `mapstructure:"repo_url"`
}

func dotlockerDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(home, ".dotlocker"), nil
}

// ConfigPath returns the path of the config file
func ConfigPath() (string, error) {
	dir, err := dotlockerDir()
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.yaml"), nil
}

// Save persists the user's configuration
func Save(repoURL string) error {
	viper.Set("repo_url", repoURL)
	return viper.WriteConfig()
}

// Load fetches the user's configuration
func Load() (Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
