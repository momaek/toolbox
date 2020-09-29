package mysql

import (
	"context"
	"time"

	log "github.com/momaek/toolbox/logger"
	"gorm.io/gorm/logger"
)

// Log implement gorm logger
type Log struct {
	log.Logger
}

// LogMode implement gorm logger
func (l *Log) LogMode(_ logger.LogLevel) logger.Interface {
	panic("not implemented") // TODO: Implement
}

// Info implement gorm logger
func (l *Log) Info(_ context.Context, _ string, _ ...interface{}) {
	panic("not implemented") // TODO: Implement
}

// Warn implement gorm logger
func (l *Log) Warn(_ context.Context, _ string, _ ...interface{}) {
	panic("not implemented") // TODO: Implement
}

// Error implement gorm logger
func (l *Log) Error(_ context.Context, _ string, _ ...interface{}) {
	panic("not implemented") // TODO: Implement
}

// Trace implement gorm logger
func (l *Log) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	panic("not implemented") // TODO: Implement
}

// newLog
func newLog(xReqID string) *Log {
	log := log.NewWithoutCaller(xReqID).Caller(7)
	return &Log{log}
}
