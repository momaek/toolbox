package cache

import (
	"time"
)

// Cacher cach interface
type Cacher interface {
	Get(interface{}) (interface{}, bool)
	Set(interface{}) error
	SetTTL(interface{}, time.Duration) error
	// If key not exist return success
	Delete(interface{})
}
