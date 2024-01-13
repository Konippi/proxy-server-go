package main

import (
	"log/slog"
	"net/http"
	"net/http/httputil"

	"github.com/Konippi/proxy-server-go/config"
	"github.com/Konippi/proxy-server-go/internal/logger"
	"github.com/Konippi/proxy-server-go/internal/path"
	"github.com/cockroachdb/errors"
)

func run(host string, port string, certFilePath string, secKeyFilePath string) {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = "localhost:8080"
	}
	handler := &httputil.ReverseProxy{Director: director}
	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: handler,
	}

	slog.Info("Server started", slog.String("host", host), slog.String("port", port))
	if err := server.ListenAndServeTLS(certFilePath, secKeyFilePath); err != nil {
		slog.Error("Server stopped", err)
	}
}

func main() {
	// initialize logger
	logger.Init()

	// initialize config
	cfg, err := config.Init()
	if err != nil {
		slog.Error("Failed to initialize config", err)
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
		slog.Error("Server started without TLS due to missing certificate or private key", tlsErr)
	}

	run(cfg.Server.Host, cfg.Server.Port, certFilePath, secKeyFilePath)
}
