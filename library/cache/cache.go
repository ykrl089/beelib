package cache

import (
	"github.com/astaxie/beego/cache"
	"time"
)

const (
	Minute = 60 * time.Second
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Month  = 30 * Day
)

var (
	MemCache cache.Cache
)

func init() {
	MemCache, _ = cache.NewCache("memory", `{"interval":60}`)
}
