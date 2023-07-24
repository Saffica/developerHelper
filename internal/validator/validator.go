package validator

import (
	"errors"
	"os"
)

type validator struct {
}

func New() *validator {
	return &validator{}
}

func (v *validator) DirExists(path string) (bool, error) {
	isExists, err := v.exists(path)
	if err != nil {
		return false, err
	}

	if !isExists {
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
