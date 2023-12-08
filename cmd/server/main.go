package main

import (
	"github.com/Konippi/proxy-server-go/config"
	"github.com/Konippi/proxy-server-go/internal/logger"
)

func main() {
	// initialize logger
	logger.Init()

	// initialize config
	config.Init()
}
