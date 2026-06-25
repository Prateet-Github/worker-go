package s3

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

func DownloadFile(
	ctx context.Context,
	client *awss3.Client,
	bucket string,
	key string,
	outputPath string,
) error {

	log.Printf("Downloading: %s", key)
	log.Printf("Saving to: %s", outputPath)

	resp, err := client.GetObject(
		ctx,
		&awss3.GetObjectInput{
			Bucket: &bucket,
			Key:    &key,
		},
	)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := os.MkdirAll(
		filepath.Dir(outputPath),
		0755,
	); err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		_ = os.Remove(outputPath)
		return err
	}

	info, err := file.Stat()
	if err != nil {
		_ = os.Remove(outputPath)
		return err
	}

	if info.Size() == 0 {
		_ = os.Remove(outputPath)
		return fmt.Errorf("downloaded file is empty")
	}

	log.Printf(
		"Downloaded: %s -> %s",
		key,
		outputPath,
	)

	return nil
}
