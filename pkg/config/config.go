package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(p string) (*viper.Viper, error) {
	cfg := viper.New()
	cfg.AutomaticEnv()
	cfg.SetConfigFile(p)
	err := cfg.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
