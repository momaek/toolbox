package session

import (
	"github.com/momaek/toolbox/utils"
)

// Session session interface
type Session interface {
	SID() string
	Get(key string) *utils.Value
	Set(key string, val interface{})
	Delete(key string)
	Exsit(key string) bool
	Destroy() error
	Clean()
	Flush() error
	Touch() error
}

// Store session store where save session data
type Store interface{}
