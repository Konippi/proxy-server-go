package config

import (
	"os"

	"github.com/Konippi/oauth-server-go/pkg/err"
	"github.com/cockroachdb/errors"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

type Config interface {
	AUTH_HOST() string
}

func init() {
	loadEnv()
}

func loadEnv() {
	envPath := "../env/.env"

	switch os.Getenv("ENV") {
	case "development":
		envPath += ".development"
	case "staging":
		envPath += ".staging"
	case "production":
		envPath += ".production"
	default:
		envPath += ".development"
	}
	godotenv.Load(envPath)
}

func AUTH_HOST() string {
	v, ok := os.LookupEnv("AUTH_HOST")
	if !ok {
		errors.Wrap(err.ErrNotFoundEnv, slog.Error("AUTH_HOST is not found in env"))
	}
	return v
}
