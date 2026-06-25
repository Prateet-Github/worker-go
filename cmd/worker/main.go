package main

import (
	"log"

	"github.com/Prateet-Github/worker-go/internal/config"
	"github.com/Prateet-Github/worker-go/internal/handlers"
	"github.com/Prateet-Github/worker-go/internal/queue"
	"github.com/Prateet-Github/worker-go/internal/s3"

	"github.com/hibiken/asynq"
)

func main() {
	cfg := config.Load()

	server := queue.NewServer(cfg)

	s3Client := s3.NewClient(cfg)
	videoHandler := handlers.NewVideoHandler(
		s3Client,
		cfg,
	)

	mux := asynq.NewServeMux()

	mux.HandleFunc(
		queue.TypeProcessVideo,
		videoHandler.ProcessVideo,
	)

	log.Println("Worker started")

	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}
}
