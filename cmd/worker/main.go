package main

import (
	"fmt"

	"github.com/Prateet-Github/worker-go/internal/config"
)

func main() {
	cfg := config.Load()
	fmt.Println("Worker Go started...")
	fmt.Printf("Redis Host: %s\n", cfg.RedisHost)
	fmt.Printf("Redis Port: %s\n", cfg.RedisPort)
}
