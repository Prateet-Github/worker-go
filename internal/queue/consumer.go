package queue

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Prateet-Github/worker-go/internal/types"
	"github.com/Prateet-Github/worker-go/internal/worker"
	"github.com/redis/go-redis/v9"
)

func ConsumeJobs(client *redis.Client) {
	ctx := context.Background()

	for {
		result, err := client.BRPop(ctx, 0, VideoQueue).Result()

		if err != nil {
			log.Println("Consumer error:", err)
			continue
		}

		var job types.VideoJob

		if err := json.Unmarshal([]byte(result[1]), &job); err != nil {
			log.Println("Invalid job:", err)
			continue
		}

		worker.Handle(job)
	}
}
