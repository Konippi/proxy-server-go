package config

import (
	"log/slog"

	"github.com/Konippi/proxy-server-go/internal/yml"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port" default:"8080"`
	} `yaml:"server"`
}

func Init() {
	f, err := yml.LoadYAML()
	if err != nil {
		slog.Error("Failed to initialize config", err)
	}
	err = yml.Deserialize(f, &Config{})
	if err != nil {
		slog.Error("Failed to initialize config", err)
	}
}
