package config

import (
	"fmt"
	"os"

	"github.com/Konippi/proxy-server-go/pkg/customerr"
	"github.com/cockroachdb/errors"
	"github.com/joho/godotenv"
)

type Config interface {
	PROXY_HOST() string
	PROXY_PORT() string
}

func init() {
	loadEnv()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		errors.Wrap(err, "Error loading env file")
	}
}

func PROXY_HOST() string {
	envName := "PROXY_HOST"
	v, ok := os.LookupEnv(envName)
	if !ok {
		errors.Wrap(customerr.ErrNotFoundEnv, fmt.Sprintf("Not Found Env: %s", envName))
	}
	return v
}

func PROXY_PORT() string {
	envName := "PROXY_PORT"
	v, ok := os.LookupEnv(envName)
	if !ok {
		errors.Wrap(customerr.ErrNotFoundEnv, fmt.Sprintf("Not Found Env: %s", envName))
	}
	return v
}
