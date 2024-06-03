package redis

import (
	ctx "context"
	"github.com/vinylSummer/microUrl/pkg/redis"
	"time"
)

type CacheRepository struct {
	*redis.Connection
}

func New(connection *redis.Connection) *CacheRepository {
	return &CacheRepository{connection}
}

func (repo *CacheRepository) Ping(context ctx.Context) error {
	return repo.Connection.Client.Ping(context).Err()
}

func (repo *CacheRepository) Clear(context ctx.Context) error {
	return repo.Connection.Client.FlushDB(context).Err()
}

func (repo *CacheRepository) Set(context ctx.Context, key string, value interface{}, expiration time.Duration) error {
	return repo.Connection.Client.Set(context, key, value, expiration).Err()
}

func (repo *CacheRepository) Get(context ctx.Context, key string) (interface{}, error) {
	return repo.Connection.Client.Get(context, key).Result()
}
