package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// Config is the regex pattern config
type Config struct {
	Pattern string `json:"pattern"`         // e.g. (?P<name>...)
	Flags   string `json:"flags,omitempty"` // e.g. "i" for case-insensitive
}

// LoadConfig reads and unmarshal the JSON config
func LoadConfig(path string, overrideConfig *Config) (*Config, error) {
	_, err := os.Stat(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("[gexv.LoadConfig] failed to check config file stat: %w", err)
	}

	cfg := new(Config)
	if !errors.Is(err, os.ErrNotExist) {
		cfg, err = Load(path)
		if err != nil {
			return nil, fmt.Errorf("[gexv.LoadConfig] failed to load config: %w", err)
		}
	}

	applyOverrides(cfg, overrideConfig)

	if err := validate(cfg); err != nil {
		return nil, fmt.Errorf("[gexv.LoadConfig] failed to validate config: %w", err)
	}

	return cfg, nil
}

func validate(cfg *Config) error {
	if cfg.Pattern == "" {
		return fmt.Errorf("pattern must be specified")
	}

	return nil
}

func applyOverrides(cfg *Config, overrideConfig *Config) {
	if overrideConfig.Pattern != "" {
		cfg.Pattern = overrideConfig.Pattern
	}

	if overrideConfig.Flags != "" {
		cfg.Flags = overrideConfig.Flags
	}
}

func Load(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
