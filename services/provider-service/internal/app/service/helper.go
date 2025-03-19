package service

import (
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func ToString(ID uuid.UUID) string {
	return ID.String()
}

func ToUUID(ID string) (uuid.UUID, error) {
	return uuid.Parse(ID)
}

func GetRootFilePath(filename string) (string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(rootDir, filename), nil
}
