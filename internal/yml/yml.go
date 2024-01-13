package yml

import (
	"os"

	"github.com/cockroachdb/errors"
	"gopkg.in/yaml.v3"
)

func Deserialize[T any](path string, target *T) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "Failed to open yml file")
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&target)
	if err != nil {
		return errors.Wrap(err, "Failed to decode yml file")
	}
	return nil
}
