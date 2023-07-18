package main

import (
	"log"

	"main/internal/config"
	"main/internal/fileservice"
	"main/internal/vcsrepository"
)

var cfg = config.GetConfig()

func main() {
	config.Validation()
	vcs, err := vcsrepository.FindByLocalPackID(cfg.GetString("LOCAL_PACK_ID"))
	if err != nil {
		log.Panic(err.Error())
	}
	fileservice.CreateFiles(vcs)
}
