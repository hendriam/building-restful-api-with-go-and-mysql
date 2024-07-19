package framework

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logging struct {
	*zerolog.Logger
}

func LoadLogging() Logging {
	config := LoadConfig()

	logger := zerolog.
		New(os.Stdout).
		Level(level(config.Log.Level)).
		With().
		Timestamp().
		Logger()

	output := logger.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	})

	return Logging{
		&output,
	}
}

func level(logLevel string) zerolog.Level {
	switch logLevel {
	case "error":
		return zerolog.ErrorLevel
	case "warn":
		return zerolog.WarnLevel
	case "info":
		return zerolog.InfoLevel
	case "debug":
		return zerolog.DebugLevel
	case "trace":
		return zerolog.TraceLevel
	default:
		/* code */
		return zerolog.InfoLevel
	}
}
