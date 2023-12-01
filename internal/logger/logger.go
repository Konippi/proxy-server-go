package logger

import (
	"log/slog"
	"os"

	"github.com/Konippi/proxy-server-go/internal/env"
)

func Initialize() {
	lebel := slog.LevelInfo
	if env.NewEnvProvider().Get() == env.Development {
		lebel = slog.LevelDebug
	}

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     lebel,
	}

	h := slog.NewJSONHandler(os.Stdout, opts)
	slog.SetDefault(slog.New(h))
}
