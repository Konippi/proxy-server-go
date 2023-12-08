package yml

import (
	"fmt"
	"os"

	"github.com/Konippi/proxy-server-go/internal/env"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func LoadYAML() (*os.File, error) {
	envStr := env.NewEnvProvider().Get().String()
	f, err := os.Open(fmt.Sprintf("../../config/yml/config.%s.yml", envStr))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open yml file")
	}
	defer f.Close()
	return f, nil
}

func Deserialize[T any](f *os.File, target T) error {
	err := yaml.NewDecoder(f).Decode(&target)
	if err != nil {
		return errors.Wrap(err, "Failed to decode target file")
	}
	return nil
}
