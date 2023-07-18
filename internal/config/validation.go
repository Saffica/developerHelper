package config

import (
	"errors"
	"log"
	"os"
)

var cfg = GetConfig()

func Validation() {
	correctPath, err := exists(cfg.GetString("TARGET_DIR_PATH"))
	if err != nil {
		log.Panic(err.Error())
	}

	if !correctPath {
		log.Panic(errors.New("bad target dir path"))
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
