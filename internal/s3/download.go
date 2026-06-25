package s3

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

func DownloadFile(
	client *awss3.Client,
	bucket string,
	key string,
	outputPath string,
) error {

	fmt.Printf("Downloading: %s\n", key)
	fmt.Printf("Saving to: %s\n", outputPath)

	resp, err := client.GetObject(
		context.Background(),
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

	fmt.Printf(
		"Downloaded: %s -> %s\n",
		key,
		outputPath,
	)

	return nil
}
