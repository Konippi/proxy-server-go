package config

import (
	"fmt"

	"github.com/Konippi/proxy-server-go/internal/env"
	"github.com/Konippi/proxy-server-go/internal/path"
	"github.com/Konippi/proxy-server-go/internal/yml"
	"github.com/cockroachdb/errors"
)

type Config struct {
	Server struct {
		Host           string `yaml:"host" default:"127.0.0.1"`
		Port           string `yaml:"port" default:"8080"`
		CertFilePath   string `yaml:"certFilePathFromRoot" default:""`
		SecKeyFilePath string `yaml:"secKeyFilePathFromRoot" default:""`
	} `yaml:"server"`
}

func Init() (Config, error) {
	envStr := env.NewEnvProvider().Get().String()
	p, err := path.LocalPath(fmt.Sprintf("config/yml/config.%s.yml", envStr))
	if err != nil {
		return Config{}, errors.Wrap(err, "Failed to initialize config")
	}

	var cfg Config
	err = yml.Deserialize(p, &cfg)
	if err != nil {
		return Config{}, errors.Wrap(err, "Failed to initialize config")
	}

	return cfg, nil
}
