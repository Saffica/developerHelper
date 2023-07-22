package config

import (
	"github.com/spf13/viper"

	"main/internal/validator"
)

type cfg struct {
	path string
}

func New(p string) *cfg {
	return &cfg{
		path: p,
	}
}

func (c *cfg) GetConfig() (*viper.Viper, error) {
	cfg, err := c.loadConfig()
	if err != nil {
		return nil, err
	}

	validator := validator.New()
	if _, err := validator.DirExists(cfg.GetString("TARGET_DIR_PATH")); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *cfg) loadConfig() (*viper.Viper, error) {
	cfg := viper.New()
	cfg.SetDefault("cfg.path", c.path)
	cfg.AutomaticEnv()
	cfg.SetConfigFile(cfg.GetString("cfg.path"))
	err := cfg.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
