package main

import (
	"log"

	"main/internal/client"
	"main/internal/config"
	"main/internal/fileservice"
	"main/internal/vcsrepository"
)

func main() {
	cfg, err := config.New().GetConfig()
	if err != nil {
		log.Panic(err.Error())
	}

	client := client.New(cfg)
	vcsrepository := vcsrepository.New(cfg, client)
	vcs, err := vcsrepository.FindByLocalPackID()
	if err != nil {
		log.Panic(err.Error())
	}
	fileservice := fileservice.New(cfg, client)
	if _, err := fileservice.CreateFiles(vcs); err != nil {
		log.Panic(err.Error())
	}

}
