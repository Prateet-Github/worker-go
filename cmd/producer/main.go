package main

import (
	"log"

	"github.com/Prateet-Github/worker-go/internal/config"
	"github.com/Prateet-Github/worker-go/internal/queue"
	"github.com/Prateet-Github/worker-go/internal/types"
)

func main() {
	cfg := config.Load()

	client := queue.NewRedisClient(cfg)
	defer client.Close()

	if err := queue.Ping(client); err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	job := types.VideoJob{
		VideoID: "12345",
		S3Key:   "videos/video1.mp4",
	}

	if err := queue.PushJob(client, job); err != nil {
		log.Fatal("Failed to push job:", err)
	}

	log.Println("Job pushed successfully")
}
