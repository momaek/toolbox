package logger

import (
	"fmt"
)

// Debug debug level
func (l *Log) Debug(arguments ...interface{}) {
	l.log.Debug().Msg(fmt.Sprint(arguments...))
}

// Info info level
func (l *Log) Info(arguments ...interface{}) {
	l.log.Info().Msg(fmt.Sprint(arguments...))
}

// Warn warn level
func (l *Log) Warn(arguments ...interface{}) {
	l.log.Warn().Msg(fmt.Sprint(arguments...))
}

// Error error level
func (l *Log) Error(arguments ...interface{}) {
	l.log.Error().Msg(fmt.Sprint(arguments...))
}

// Fatal fatal level
func (l *Log) Fatal(arguments ...interface{}) {
	l.log.Fatal().Msg(fmt.Sprint(arguments...))
}

// Panic panic level
func (l *Log) Panic(arguments ...interface{}) {
	l.log.Panic().Msg(fmt.Sprint(arguments...))
}

// Debugf debug format
func (l *Log) Debugf(format string, arguments ...interface{}) {
	l.log.Debug().Msg(fmt.Sprintf(format, arguments...))
}

// Infof info format
func (l *Log) Infof(format string, arguments ...interface{}) {
	l.log.Info().Msg(fmt.Sprintf(format, arguments...))
}

// Warnf warn format
func (l *Log) Warnf(format string, arguments ...interface{}) {
	l.log.Warn().Msg(fmt.Sprintf(format, arguments...))
}

// Errorf error format
func (l *Log) Errorf(format string, arguments ...interface{}) {
	l.log.Error().Msg(fmt.Sprintf(format, arguments...))
}

// Fatalf fatal format
func (l *Log) Fatalf(format string, arguments ...interface{}) {
	l.log.Fatal().Msg(fmt.Sprintf(format, arguments...))
}

// Panicf panic format
func (l *Log) Panicf(format string, arguments ...interface{}) {
	l.log.Panic().Msg(fmt.Sprintf(format, arguments...))
}

// ReqID http rpc call
func (l *Log) ReqID() string {
	return l.reqID
}

// Xput rpc call may use
func (l *Log) Xput(logs []string) {
	panic("not implemented") // TODO: Implement
}
