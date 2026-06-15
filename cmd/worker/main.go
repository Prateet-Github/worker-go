package main

import (
	"fmt"

	"github.com/Prateet-Github/worker-go/internal/config"
	"github.com/Prateet-Github/worker-go/internal/queue"
	"github.com/Prateet-Github/worker-go/internal/types"
)

func main() {
	cfg := config.Load()
	client := queue.NewRedisClient(cfg)

	fmt.Println("Worker Go started...")
	fmt.Printf("Redis Host: %s\n", cfg.RedisHost)
	fmt.Printf("Redis Port: %s\n", cfg.RedisPort)

	err := queue.Ping(client)
	if err != nil {
		fmt.Printf("Error connecting to Redis: %v\n", err)
		return
	}

	fmt.Println("Connected to Redis!")

	err = queue.PushJob(client, types.VideoJob{
		VideoID: "12345",
		S3Key:   "videos/video1.mp4",
	})

	if err != nil {
		fmt.Printf("Error pushing job: %v\n", err)
		return
	}

	fmt.Println("Job pushed to queue!")

	queue.ConsumeJobs(client)
}
