package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisHost string
	RedisPort string

	AWSRegion          string
	AWSAccessKeyID     string
	AWSSecretAccessKey string

	S3RawBucket  string
	S3ProdBucket string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		RedisHost:          os.Getenv("REDIS_HOST"),
		RedisPort:          os.Getenv("REDIS_PORT"),
		AWSRegion:          os.Getenv("AWS_REGION"),
		AWSAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		S3RawBucket:        os.Getenv("S3_RAW_BUCKET"),
		S3ProdBucket:       os.Getenv("S3_PROD_BUCKET"),
	}
}
