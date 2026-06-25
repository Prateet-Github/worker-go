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

	APIBaseURL string
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using environment variables")
	}

	return &Config{
		RedisHost:          os.Getenv("REDIS_HOST"),
		RedisPort:          os.Getenv("REDIS_PORT"),
		AWSRegion:          os.Getenv("AWS_REGION"),
		AWSAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		S3RawBucket:        os.Getenv("S3_RAW_BUCKET"),
		S3ProdBucket:       os.Getenv("S3_PROD_BUCKET"),
		APIBaseURL:         os.Getenv("API_BASE_URL"),
	}
}
