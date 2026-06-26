package logging

import (
	"github.com/dElCIoGio/filestorage/internal/platform/config"
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	cfg := config.GetConfig()

	if cfg.Env == config.PROD {
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}
