package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisHost string
	RedisPort string
	AWSRegion string
	RawBucket string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
		AWSRegion: os.Getenv("AWS_REGION"),
		RawBucket: os.Getenv("S3_RAW_BUCKET"),
	}
}
