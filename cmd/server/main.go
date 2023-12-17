package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/Konippi/proxy-server-go/config"
	"github.com/Konippi/proxy-server-go/internal/logger"
)

func main() {
	// initialize logger
	logger.Init()

	// initialize config
	cfg, err := config.Init()
	if err != nil {
		slog.Error("Failed to initialize config", err)
	}

	// start server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, r.Host) })
	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		ReadHeaderTimeout: 20 * time.Second,
	}
	slog.Info("Server started", slog.String("host", cfg.Server.Host), slog.String("port", cfg.Server.Port))
	server.ListenAndServeTLS(cfg.Server.CertFilePath, cfg.Server.SecKeyFilePath)
}
