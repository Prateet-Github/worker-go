package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DeleteObject(
	ctx context.Context,
	client *s3.Client,
	bucket string,
	key string,
) error {
	_, err := client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})

	if err != nil {
		return err
	}

	return nil
}
