package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

// New returns an instntiated, configured logger
func New(logLevel string, useConsole bool, fields map[string]interface{}) zerolog.Logger {

	var w io.Writer = os.Stdout

	if useConsole {
		w = zerolog.ConsoleWriter{
			Out: os.Stdout,
		}
	}

	lvl, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		lvl = zerolog.InfoLevel
	}

	return zerolog.New(w).Level(lvl).With().Fields(fields).Timestamp().Logger()
}
