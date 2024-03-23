package logger

import (
	"context"
	"log/slog"
)

var globalLogger *slog.Logger

func Init(logger *slog.Logger) {
	globalLogger = logger
}

func Debug(msg string, fields ...slog.Attr) {
	globalLogger.LogAttrs(context.Background(), slog.LevelDebug, msg, fields...)
}

func Info(msg string, fields ...slog.Attr) {
	globalLogger.LogAttrs(context.Background(), slog.LevelInfo, msg, fields...)
}

func Warn(msg string, fields ...slog.Attr) {
	globalLogger.LogAttrs(context.Background(), slog.LevelWarn, msg, fields...)
}

func Error(msg string, fields ...slog.Attr) {
	globalLogger.LogAttrs(context.Background(), slog.LevelError, msg, fields...)
}
