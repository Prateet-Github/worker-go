package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Prateet-Github/worker-go/internal/queue"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hibiken/asynq"
)

type VideoHandler struct {
	s3Client *s3.Client
}

func NewVideoHandler(
	s3Client *s3.Client,
) *VideoHandler {
	return &VideoHandler{
		s3Client: s3Client,
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

	log.Println("Processing video...")
	log.Println("Video ID:", payload.VideoID)
	log.Println("S3 Key:", payload.S3Key)

	return nil
}
