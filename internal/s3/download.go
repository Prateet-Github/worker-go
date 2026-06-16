package s3

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DownloadFile(
	client *s3.Client,
	bucket string,
	key string,
	localPath string,
) error {

	resp, err := client.GetObject(
		context.Background(),
		&s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &key,
		},
	)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	file, err := os.Create(localPath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	return err
}
