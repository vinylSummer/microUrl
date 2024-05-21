package repositories

import (
	ctx "context"
	"time"
)

type CacheRepository interface {
	Ping(context ctx.Context) error
	Clear(context ctx.Context) error
	Set(context ctx.Context, key string, value interface{}, expiration time.Duration) error
	Get(context ctx.Context, key string) (interface{}, error)
}
