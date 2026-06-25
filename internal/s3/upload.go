package s3

import (
	"context"
	"log"
	"mime"
	"os"
	"path/filepath"

	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadDirectory(
	ctx context.Context,
	client *awss3.Client,
	bucket string,
	localDir string,
	s3Prefix string,
) error {

	return filepath.Walk(
		localDir,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			relativePath, err := filepath.Rel(localDir, path)
			if err != nil {
				return err
			}

			s3Key := filepath.ToSlash(
				filepath.Join(s3Prefix, relativePath),
			)

			if err := uploadFile(
				ctx,
				client,
				bucket,
				path,
				s3Key,
			); err != nil {
				return err
			}

			log.Printf(
				"Uploaded: %s -> %s",
				path,
				s3Key,
			)

			return nil
		},
	)
}

func uploadFile(
	ctx context.Context,
	client *awss3.Client,
	bucket string,
	localPath string,
	s3Key string,
) error {

	file, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	contentType := mime.TypeByExtension(
		filepath.Ext(localPath),
	)

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err = client.PutObject(
		ctx,
		&awss3.PutObjectInput{
			Bucket:      &bucket,
			Key:         &s3Key,
			Body:        file,
			ContentType: &contentType,
		},
	)

	return err
}
