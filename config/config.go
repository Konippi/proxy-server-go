package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"github.com/cockroachdb/errors"
	"github.com/joho/godotenv"
)

type Config struct {
	ProxyHost string
	ProxyPort string
}

func init() {
	loadEnv()
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return errors.Wrap(err, "Error loading env file")
	}
	slog.Info("Successfully loaded env file")
	return nil
}

func (c *Config) Dump(w io.Writer) error {
	fmt.Fprint(w, "config: ")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(c)
}
