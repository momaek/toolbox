package localcache

import (
	"github.com/allegro/bigcache"
)

func new() {
	c := bigcache.NewBigCache(bigcache.Config{})
}
