package queue

import (
	"context"
	"encoding/json"

	"github.com/Prateet-Github/worker-go/internal/types"
	"github.com/redis/go-redis/v9"
)

func PushJob(client *redis.Client, job types.VideoJob) error {
	ctx := context.Background()

	payload, err := json.Marshal(job)

	if err != nil {
		return err
	}

	return client.LPush(ctx, VideoQueue, payload).Err()
}
