package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	l := NewWithCaller()

	l.Info("hello")

	l.WithField(map[string]interface{}{"type": "DATABASE"}).Info("helllo3")

	l.SetLevel(Warn).Warn("hello")
}
