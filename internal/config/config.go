package config

import (
	"log"
	"main/pkg/config"

	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
	cfg, err := config.LoadConfig("./configs/config.json")
	if err != nil {
		log.Panic(err.Error())
	}
	return cfg
}
