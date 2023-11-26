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
	if v, ok := os.LookupEnv(envName); ok {
		return v
	}
	errors.Wrap(customerr.ErrNotFoundEnv, fmt.Sprintf("Not Found Env: %s", envName))

	return "127.0.0.1"
}
