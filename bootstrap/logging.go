package bootstrap

import (
	"log/slog"
	"os"
	"strings"

	"github.com/uptrace/mcp/appconf"
	"go.uber.org/fx"
)

type LoggerResults struct {
	fx.Out

	Logger *slog.Logger
}

func NewSlog(conf *appconf.Config) LoggerResults {
	level := new(slog.LevelVar)

	switch strings.ToLower(conf.Logging.Level) {
	case "debug":
		level.Set(slog.LevelDebug)
	case "info":
		level.Set(slog.LevelInfo)
	case "warn":
		level.Set(slog.LevelWarn)
	case "error":
		level.Set(slog.LevelError)
	default:
		level.Set(slog.LevelInfo)
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	}))

	return LoggerResults{
		Logger: logger,
	}
}
