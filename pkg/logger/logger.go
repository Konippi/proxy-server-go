package logger

import (
	"log/slog"
	"os"

	"github.com/Konippi/proxy-server-go/pkg/env"
)

func Init() {
	level := slog.LevelInfo
	if env.NewEnvProvider().Get() == env.Development {
		level = slog.LevelDebug
	}

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}

	h := slog.NewJSONHandler(os.Stdout, opts)
	slog.SetDefault(slog.New(h))
}
