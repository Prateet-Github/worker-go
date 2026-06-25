package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Prateet-Github/worker-go/internal/config"
	"github.com/Prateet-Github/worker-go/internal/queue"
	"github.com/Prateet-Github/worker-go/internal/s3"

	"path/filepath"

	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hibiken/asynq"
)

type VideoHandler struct {
	s3Client *awss3.Client
	cfg      *config.Config
}

func NewVideoHandler(
	s3Client *awss3.Client,
	cfg *config.Config,
) *VideoHandler {
	return &VideoHandler{
		s3Client: s3Client,
		cfg:      cfg,
	}
}

func (h *VideoHandler) ProcessVideo(
	ctx context.Context,
	task *asynq.Task,
) error {

	var payload queue.VideoTask

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	localPath := filepath.Join(
		"temp",
		payload.VideoID+".mp4",
	)

	err := s3.DownloadFile(
		h.s3Client,
		h.cfg.S3RawBucket,
		payload.S3Key,
		localPath,
	)

	if err != nil {
		return err
	}

	log.Println("Download completed:", localPath)

	log.Println("Processing video...")
	log.Println("Video ID:", payload.VideoID)
	log.Println("S3 Key:", payload.S3Key)

	return nil
}
