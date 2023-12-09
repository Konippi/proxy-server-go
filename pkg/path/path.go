package path

import (
	"path/filepath"
	"runtime"

	"github.com/cockroachdb/errors"
)

func RootDir() (string, error) {
	_, currentPath, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("Failed to get current path")
	}
	currentDir := filepath.Dir(currentPath)
	p := filepath.Join(currentDir, "..", "..")
	return p, nil
}

func CallerPath() (string, error) {
	_, p, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("Failed to get caller path")
	}
	return p, nil
}
