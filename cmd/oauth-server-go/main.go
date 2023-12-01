package main

import (
	"os"

	"github.com/Konippi/proxy-server-go/config/serverconfig"
	"github.com/cockroachdb/errors"
)

func main() {
	cfg, err := serverconfig.Init()
	if err != nil {
		errors.Wrap(err, "Failed to initialize server config")
	}
	cfg.Dump(os.Stdout)

	// get environment variables
	// SERVER_PORT := os.Getenv("BACKEND_PORT")
}
