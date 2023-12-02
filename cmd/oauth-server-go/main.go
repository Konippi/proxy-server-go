package main

import (
	"os"

	"github.com/Konippi/proxy-server-go/config/serverconfig"
	"github.com/Konippi/proxy-server-go/internal/logger"
	"github.com/cockroachdb/errors"
)

func main() {
	// initialize logger
	logger.Init()

	cfg, err := serverconfig.Init()
	if err != nil {
		errors.Wrap(err, "Failed to initialize server config")
	}
	cfg.Dump(os.Stdout)

	// get environment variables
	// SERVER_PORT := os.Getenv("BACKEND_PORT")
}
