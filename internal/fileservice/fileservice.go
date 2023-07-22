package fileservice

import (
	"fmt"
	"os"

	"main/internal/model"
)

type Config interface {
	GetString(string) string
}

type Client interface {
	GetRecordsByVCS(vcs map[string]model.Table) ([]model.Record, error)
}

type app struct {
	client Client
	cfg    Config
}

func New(cfg Config, client Client) *app {
	return &app{
		client: client,
		cfg:    cfg,
	}
}

func (a *app) CreateFiles(recordCandidates map[string]model.Table) (bool, error) {
	files, err := a.client.GetRecordsByVCS(recordCandidates)
	if err != nil {
		return false, err
	}

	for _, i := range files {
		currentPath := fmt.Sprintf("%s\\%s", a.cfg.GetString("TARGET_DIR_PATH"), i.TableName)
		createDir(currentPath)

		if i.TableName == "sys_widget" {
			currentPath = fmt.Sprintf("%s\\%s", currentPath, i.SysID)
			createDir(currentPath)
		}

		file := fmt.Sprintf("%s\\%s", currentPath, i.FileName)
		f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			return false, err
		}

		defer f.Close()

		f.WriteString(i.Value)
	}

	return true, nil
}

<<<<<<< HEAD
func createDir(p string) (bool, error) {
	folderIsExist, err := exists(p)
	if err != nil {
		return false, err
=======
func createDir(p string) {
	folderIsExist, err := exists(p)
	if err != nil {
		log.Panic(err.Error())
>>>>>>> master
	}

	if !folderIsExist {
		err = os.Mkdir(p, 0777)
		if err != nil {
<<<<<<< HEAD
			return false, err
		}
	}

	return true, nil
=======
			log.Panic(err.Error())
		}
	}
>>>>>>> master
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
