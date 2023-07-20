package config

import (
	"github.com/spf13/viper"
)

type cfg struct {
}

func New() *cfg {
	return &cfg{}
}

func (c *cfg) GetConfig() (*viper.Viper, error) {
	cfg, err := c.loadConfig()
	if err != nil {
		return nil, err
	}

	validator := NewValidator(cfg)
	if _, err := validator.Run(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *cfg) loadConfig() (*viper.Viper, error) {
	cfg := viper.New()
	cfg.SetDefault("cfg.path", "./configs/config.json")
	cfg.AutomaticEnv()
	cfg.SetConfigFile(cfg.GetString("cfg.path"))
	err := cfg.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
