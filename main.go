package main

import (
	"fmt"
	"log"
	"os"

	"main/internal/client"
	"main/internal/config"
)

func main() {
	cfg := config.GetConfig()
	correctPath, err := client.Exists(cfg.GetString("TARGET_DIR_PATH"))
	if err != nil {
		log.Panic(err.Error())
	}

	if !correctPath {
		log.Panic("Incorrect TARGET_DIR_PATH")
	}

	vcs, err := client.GetVCS()
	if err != nil {
		log.Panic(err.Error())
	}

	vcsForRequest, err := client.PrepareVCS(vcs)
	if err != nil {
		log.Panic(err.Error())
	}

	files, err := client.GetRecordsByVCS(vcsForRequest)
	if err != nil {
		log.Panic(err.Error())
	}

	for _, i := range files {
		dir := fmt.Sprintf("%s\\%s", cfg.GetString("TARGET_DIR_PATH"), i.TableName)
		folderIsExist, err := client.Exists(dir)
		if err != nil {
			log.Panic(err.Error())
		}

		if !folderIsExist {
			err = os.Mkdir(dir, 0777)
			if err != nil {
				log.Panic(err.Error())
			}
		}

		file := fmt.Sprintf("%s\\%s", dir, i.FileName)
		f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			log.Panic(err.Error())
		}

		defer f.Close()

		f.WriteString(i.Value)
	}
}
