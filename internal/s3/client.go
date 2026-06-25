package s3

import (
	"context"
	"log"

	"github.com/Prateet-Github/worker-go/internal/config"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewClient(cfg *config.Config) *awss3.Client {
	awsCfg, err := awsconfig.LoadDefaultConfig(
		context.Background(),
		awsconfig.WithRegion(cfg.AWSRegion),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AWSAccessKeyID,
				cfg.AWSSecretAccessKey,
				"",
			),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	return awss3.NewFromConfig(awsCfg)
}
