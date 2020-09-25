package logger

import (
	"github.com/rs/zerolog"
)

// Level log level
type Level string

const (
	// Trace level
	Trace Level = "trace"
	// Debug level
	Debug Level = "debug"
	// Warn level
	Warn Level = "warn"
	// Info level
	Info Level = "info"
	// Error level
	Error Level = "error"
)

// ToZeroLogLevel to zerolog.Level
func (l Level) ToZeroLogLevel() zerolog.Level {
	switch l {
	case Debug:
		return zerolog.DebugLevel
	case Warn:
		return zerolog.WarnLevel
	case Info:
		return zerolog.InfoLevel
	case Error:
		return zerolog.ErrorLevel
	default:
		return zerolog.TraceLevel
	}
}
