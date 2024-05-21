package redis

import (
	"github.com/redis/go-redis/v9"
	cfg "github.com/vinylSummer/microUrl/config"
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

	return &Connection{
		Client: newClient,
	}, nil
}

func (redisClient *Connection) Close() {
	if redisClient.Client != nil {
		redisClient.Client.Close()
	}
}
