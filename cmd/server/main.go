package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Konippi/proxy-server-go/config/serverconfig"
	"github.com/Konippi/proxy-server-go/internal/logger"
	"github.com/cockroachdb/errors"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Header)

	writer.WriteHeader(200)
}

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
	server := &http.Server{
		Handler: NewHandler(),
		Addr:    "localhost:8080",
	}
	server.ListenAndServe()
}
