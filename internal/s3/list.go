package s3

import (
	"context"

	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

func ListObjects(client *awss3.Client, bucket string) error {
	maxKeys := int32(10)

	result, err := client.ListObjectsV2(
		context.Background(),
		&awss3.ListObjectsV2Input{
			Bucket:  &bucket,
			MaxKeys: &maxKeys,
		},
	)

	if err != nil {
		return err
	}

	for _, obj := range result.Contents {
		println(*obj.Key)
	}

	return nil
}
