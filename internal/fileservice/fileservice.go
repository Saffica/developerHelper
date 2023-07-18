package fileservice

import (
	"fmt"
	"log"
	"os"

	"main/internal/client"
	"main/internal/config"
)

var cfg = config.GetConfig()

func CreateFiles(recordCandidates map[string]client.Table) (bool, error) {
	files, err := client.GetRecordsByVCS(recordCandidates)
	if err != nil {
		log.Panic(err.Error())
	}

	for _, i := range files {
		dir := fmt.Sprintf("%s\\%s", cfg.GetString("TARGET_DIR_PATH"), i.TableName)
		folderIsExist, err := exists(dir)
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

	return true, nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
