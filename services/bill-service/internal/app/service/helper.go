package service

import (
	"os"
	"path/filepath"
)

func GetRootFilePath(filename string) (string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(rootDir, filename), nil
}
