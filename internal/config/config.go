// Package config provides access to the CLI's configuration
package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	RepoURL string `mapstructure:"repo_url"`
	RepoPath string `mapstructure:"repo_path"`
}

// dotlockerDir creates and returns the directory containing Dotlocker's data
func dotlockerDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(home, ".dotlocker")

	if err := os.MkdirAll(path, 0o755); err != nil {
		return "", err
	}

	return path, nil
}

// ConfigPath returns the path of the config file
func ConfigPath() (string, error) {
	dir, err := dotlockerDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.yaml"), nil
}

// SaveConfig writes an entire config to disk
func SaveConfig(v *viper.Viper, cfg Config) error {
	v.Set("repo_url", cfg.RepoURL)
	v.Set("repo_path", cfg.RepoPath)
	return viper.WriteConfig()
}

// Load fetches the user's configuration
func Load(v *viper.Viper) (Config, error) {
	if err := v.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
