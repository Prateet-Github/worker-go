package main

import (
	"log"

	"github.com/Prateet-Github/worker-go/internal/config"
	"github.com/Prateet-Github/worker-go/internal/queue"
	"github.com/Prateet-Github/worker-go/internal/s3"
	"github.com/Prateet-Github/worker-go/internal/types"
)

func main() {
	cfg := config.Load()

	// Redis
	redisClient := queue.NewRedisClient(cfg)
	defer redisClient.Close()

	if err := queue.Ping(redisClient); err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	// S3
	s3Client, err := s3.NewClient(cfg)
	if err != nil {
		log.Fatal("Failed to create S3 client:", err)
	}

	log.Println("S3 client created successfully")

	if err := s3.ListObjects(s3Client, cfg.S3RawBucket); err != nil {
		log.Fatal("Failed to list objects:", err)
	}

	_ = s3Client

	job := types.VideoJob{
		VideoID: "12345",
		S3Key:   "videos/video1.mp4",
	}

	if err := queue.PushJob(redisClient, job); err != nil {
		log.Fatal("Failed to push job:", err)
	}

	log.Println("Job pushed successfully")
}
