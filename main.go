package main

import (
	// "log"

	"main/internal/client"
	// "main/pkg/config"
)

func main() {
	// const pathToConfig = "./configs/config.json"
	// cfg, err := config.LoadConfig(pathToConfig)
	// if err != nil {
	// 	log.Panic(err.Error())
	// }

	client.GetTargetTables()
	// fmt.Println(cfg.GetString("INSTANCE_URL"))
}
