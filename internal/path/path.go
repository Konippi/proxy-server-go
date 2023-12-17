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
	d := filepath.Join(currentDir, "..", "..")
	return d, nil
}

func LocalPath(relPathFromRoot string) (string, error) {
	relPathFromRoot = filepath.Clean(relPathFromRoot)
	rootDir, err := RootDir()
	if err != nil {
		return "", errors.Wrap(err, "Failed to get root dir")
	}
	p := filepath.Join(rootDir, relPathFromRoot)
	return p, nil
}
