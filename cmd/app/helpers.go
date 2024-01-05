package app

import (
	"log/slog"
	"os"

	"github.com/Mubasshira-Tazrin/rate-limiting-api/internal/slogconf"
)

func initSlogger() *slog.Logger {
	handlerOpts := slogconf.GetSlogConf()
	l := slog.New(slog.NewTextHandler(os.Stdout, handlerOpts))
	slog.SetDefault(l)

	return l
}
