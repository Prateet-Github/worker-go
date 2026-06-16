package processor

import (
	"os"
	"path/filepath"
)

func CreateWorkspace(videoID string) (string, error) {
	path := filepath.Join("temp", videoID)

	err := os.MkdirAll(path, os.ModePerm)

	return path, err
}
