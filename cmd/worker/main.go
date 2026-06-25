package main

import (
	"log"

	"github.com/Prateet-Github/worker-go/internal/config"
	"github.com/Prateet-Github/worker-go/internal/queue"
)

func main() {
	cfg := config.Load()

	client := queue.NewClient(cfg)
	defer client.Close()

	if err := queue.Ping(client); err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	log.Println("Connected to Redis")
	log.Println("Worker started")
}
