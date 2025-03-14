package nexus

import (
	"errors"
	"os"
	"path/filepath"
)

func ReadFile(path string) (string, error) {
	fileExtension := filepath.Ext(path)
	if fileExtension != ".neo" {
		return "", errors.New("File is not a Neo file.")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
