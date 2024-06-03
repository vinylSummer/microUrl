package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	cfg "github.com/vinylSummer/microUrl/config"
	"time"
)

type Connection struct {
	Client *redis.Client
}

func New(config *cfg.Config) (*Connection, error) {
	connectionOptions, err := redis.ParseURL(config.Redis.URL)
	if err != nil {
		return nil, err
	}

	newClient := redis.NewClient(connectionOptions)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = newClient.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return &Connection{
		Client: newClient,
	}, nil
}

func (redisClient *Connection) Close() {
	if redisClient.Client != nil {
		redisClient.Client.Close()
	}
}
