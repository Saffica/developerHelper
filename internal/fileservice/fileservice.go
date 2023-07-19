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
		currentPath := fmt.Sprintf("%s\\%s", cfg.GetString("TARGET_DIR_PATH"), i.TableName)
		createDir(currentPath)

		if i.TableName == "sys_widget" {
			currentPath = fmt.Sprintf("%s\\%s", currentPath, i.SysID)
			createDir(currentPath)
		}

		file := fmt.Sprintf("%s\\%s", currentPath, i.FileName)
		f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			log.Panic(err.Error())
		}

		defer f.Close()

		f.WriteString(i.Value)
	}

	return true, nil
}

func createDir(p string) {
	folderIsExist, err := exists(p)
	if err != nil {
		log.Panic(err.Error())
	}

	if !folderIsExist {
		err = os.Mkdir(p, 0777)
		if err != nil {
			log.Panic(err.Error())
		}
	}
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
