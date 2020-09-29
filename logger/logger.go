package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

// Logger logger methods
type Logger interface {
	// STD log
	Debug(arguments ...interface{})
	Info(arguments ...interface{})
	Warn(arguments ...interface{})
	Error(arguments ...interface{})
	Fatal(arguments ...interface{})
	Panic(arguments ...interface{})
	Debugf(format string, arguments ...interface{})
	Infof(format string, arguments ...interface{})
	Warnf(format string, arguments ...interface{})
	Errorf(format string, arguments ...interface{})
	Fatalf(format string, arguments ...interface{})
	Panicf(format string, arguments ...interface{})

	// ReqID, http rpc call
	ReqID() string
	Xput(logs []string)

	// Option
	LOptioner
}

// LOptioner ..
type LOptioner interface {
	SetLevel(level Level) Logger
	WithField(field map[string]interface{}) Logger
	Caller(frame int) Logger
}

// Log implement Logger, RequestIDer, XLogger
type Log struct {
	log   *zerolog.Logger
	reqID string
}

const (
	xreqidField          = "x-reqid"
	defaultLogTimeFormat = "2006-01-02 15:04:05.000000"
)

var (
	logSyncPool = sync.Pool{New: func() interface{} {
		return ""
	}}
)

// New logger default writer is os.StdErr
// Default field: x-reqid, time, caller, message
func New(reqID ...string) Logger {
	log := newLogger(reqID...)
	l := log.(*Log)
	zl := l.log.With().CallerWithSkipFrameCount(3).Logger()
	l.log = &zl
	return l

}

// NewWithoutCaller new log without caller field
func NewWithoutCaller(reqID ...string) Logger {
	return newLogger(reqID...)
}

// newLogger return logger without caller field
func newLogger(reqID ...string) Logger {
	reqid := GenReqID()
	if len(reqID) > 0 {
		reqid = reqID[0]
	}

	l := zerolog.New(os.Stderr).With().Str(xreqidField, reqid).Timestamp().Logger()
	log := &Log{log: &l, reqID: reqid}
	SetTimeFieldFormat(defaultLogTimeFormat)
	return log
}

// WithField add new field
func (l *Log) WithField(field map[string]interface{}) Logger {
	zl := l.log.With().Fields(field).Logger()

	return &Log{
		log:   &zl,
		reqID: l.reqID,
	}
}

// SetLevel set level
func (l *Log) SetLevel(level Level) Logger {
	zl := l.log.Level(level.ToZeroLogLevel())

	return &Log{
		log:   &zl,
		reqID: l.reqID,
	}
}

// Caller set caller frame
func (l *Log) Caller(frame int) Logger {
	zl := l.log.With().CallerWithSkipFrameCount(frame).Logger()
	return &Log{
		log:   &zl,
		reqID: l.reqID,
	}
}

// SetTimeFieldFormat ..
func SetTimeFieldFormat(format string) {
	zerolog.TimeFieldFormat = format
}
