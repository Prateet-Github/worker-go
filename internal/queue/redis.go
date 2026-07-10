package queue

import (
	"context"
	"fmt"

	"github.com/Prateet-Github/worker-go/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewClient(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
	})
}

func Ping(client *redis.Client) error {
	ctx := context.Background()
	return client.Ping(ctx).Err()
}
