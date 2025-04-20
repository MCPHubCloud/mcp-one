package utils

import (
	"log/slog"
	"os"
)

// Logger 全局日志记录器
var Logger = func() *slog.Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	return slog.New(handler)
}()
