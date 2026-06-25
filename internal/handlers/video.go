package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Prateet-Github/worker-go/internal/config"
	"github.com/Prateet-Github/worker-go/internal/ffmpeg"
	"github.com/Prateet-Github/worker-go/internal/queue"
	"github.com/Prateet-Github/worker-go/internal/s3"

	"path/filepath"

	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hibiken/asynq"
)

type VideoHandler struct {
	s3Client *awss3.Client
	cfg      *config.Config
	ffmpeg   *ffmpeg.Service
}

func NewVideoHandler(
	s3Client *awss3.Client,
	cfg *config.Config,
	ffmpeg *ffmpeg.Service,
) *VideoHandler {
	return &VideoHandler{
		s3Client: s3Client,
		cfg:      cfg,
		ffmpeg:   ffmpeg,
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

	// localPath := filepath.Join(
	// 	"temp",
	// 	payload.VideoID+".mp4",
	// )

	workspace := filepath.Join(
		"temp",
		payload.VideoID,
	)

	inputPath := filepath.Join(
		workspace,
		"input.mp4",
	)

	outputDir := filepath.Join(
		workspace,
		"hls",
	)

	log.Println("Starting download...")

	err := s3.DownloadFile(
		h.s3Client,
		h.cfg.S3RawBucket,
		payload.S3Key,
		inputPath,
	)

	if err != nil {
		return err
	}

	log.Println("Download complete")

	log.Println("Calling FFmpeg...")

	if err := h.ffmpeg.GenerateHLS(
		ctx,
		inputPath,
		outputDir,
	); err != nil {
		log.Printf("GenerateHLS failed: %v\n", err)
		return err
	}

	err = s3.UploadDirectory(
		h.s3Client,
		h.cfg.S3ProdBucket,
		outputDir,
		filepath.Join(s3.ProcessedVideosPrefix, payload.VideoID),
	)

	if err != nil {
		return err
	}

	log.Println("HLS uploaded successfully")

	log.Println("FFmpeg finished")

	log.Println("HLS generated successfully")

	log.Println("Download completed:", inputPath)

	log.Println("Processing video...")
	log.Println("Video ID:", payload.VideoID)
	log.Println("S3 Key:", payload.S3Key)

	return nil
}
