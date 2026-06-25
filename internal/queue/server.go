package queue

import (
	"github.com/Prateet-Github/worker-go/internal/config"
	"github.com/hibiken/asynq"
)

func NewServer(cfg *config.Config) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: cfg.RedisHost + ":" + cfg.RedisPort,
		},
		asynq.Config{
			Concurrency: 5,
		},
	)
}
