package logger

import (
	"os"

	"github.com/Konippi/oauth-server-go/pkg/env"
	"golang.org/x/exp/slog"
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
