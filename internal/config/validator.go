package config

import (
	"errors"
	"os"
)

type Config interface {
	GetString(string) string
}

type validator struct {
	cfg Config
}

func NewValidator(cfg Config) *validator {
	return &validator{
		cfg: cfg,
	}
}

func (v *validator) Run() (bool, error) {
	correctPath, err := v.exists(v.cfg.GetString("TARGET_DIR_PATH"))
	if err != nil {
		return false, err
	}

	if !correctPath {
		return false, errors.New("bad target dir path")
	}

	return true, nil
}

func (v *validator) exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
