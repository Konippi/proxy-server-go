package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/Konippi/proxy-server-go/config"
	"github.com/Konippi/proxy-server-go/internal/logger"
	"github.com/Konippi/proxy-server-go/internal/path"
	"github.com/cockroachdb/errors"
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

	var tlsErr error
	certFilePath, err := path.LocalPath(cfg.Server.CertFilePath)
	if err != nil {
		tlsErr = errors.Wrap(err, "Failed to get certificate path")
	}
	secKeyFilePath, err := path.LocalPath(cfg.Server.SecKeyFilePath)
	if err != nil {
		errors.Wrap(tlsErr, "Failed to get private key path")
	}

	if tlsErr != nil {
		slog.Error("Server started without TLS due to missing certificate or private key")
		server.ListenAndServe()
	}

	slog.Info("Server started", slog.String("host", cfg.Server.Host), slog.String("port", cfg.Server.Port))
	if err := server.ListenAndServeTLS(certFilePath, secKeyFilePath); err != nil {
		slog.Error("Server stopped", err)
	}
}
