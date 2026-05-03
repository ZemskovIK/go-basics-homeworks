package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(path string) ([]byte, error) {
	ext := strings.ToLower(filepath.Ext(path))
	if ext != ".json" {
		return nil, errors.New("file extension is not .json")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}
