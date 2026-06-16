package s3

import (
	"context"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewClient() (*s3.Client, error) {
	cfg, err := awsconfig.LoadDefaultConfig(context.Background())

	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}
