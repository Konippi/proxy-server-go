package config

import (
	"fmt"
	"os"

	"github.com/Konippi/oauth-server-go/pkg/customerr"
	"github.com/cockroachdb/errors"
	"github.com/joho/godotenv"
)

type Config interface {
	AUTH_HOST() string
	AUTH_PORT() string
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
	envName := "AUTH_HOST"
	v, ok := os.LookupEnv(envName)
	if !ok {
		errors.Wrap(customerr.ErrNotFoundEnv, fmt.Sprintf("Not Found Env: %s", envName))
	}
	return v
}

func AUTH_PORT() string {
	envName := "AUTH_PORT"
	v, ok := os.LookupEnv(envName)
	if !ok {
		errors.Wrap(customerr.ErrNotFoundEnv, fmt.Sprintf("Not Found Env: %s", envName))
	}
	return v
}
