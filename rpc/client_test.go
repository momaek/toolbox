package rpc

import (
	"testing"

	"github.com/momaek/toolbox/logger"
)

func TestGet(t *testing.T) {
	c := New()
	_, _ = c.Get(logger.New("testclient"), "https://163.com")
}
