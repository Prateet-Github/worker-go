package handlers

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/Prateet-Github/worker-go/internal/api"
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
	api      *api.Client
}

func NewVideoHandler(
	s3Client *awss3.Client,
	cfg *config.Config,
	ffmpeg *ffmpeg.Service,
	api *api.Client,
) *VideoHandler {
	return &VideoHandler{
		s3Client: s3Client,
		cfg:      cfg,
		ffmpeg:   ffmpeg,
		api:      api,
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

	workspace := filepath.Join(
		"temp",
		payload.VideoID,
	)

	defer func() {
		if err := os.RemoveAll(workspace); err != nil {
			log.Printf("Failed to cleanup workspace %s: %v", workspace, err)
		} else {
			log.Printf("Workspace cleaned: %s", workspace)
		}
	}()

	inputPath := filepath.Join(
		workspace,
		"input.mp4",
	)

	outputDir := filepath.Join(
		workspace,
		"hls",
	)

	log.Printf("Processing video: %s", payload.VideoID)

	log.Println("Downloading video...")

	if err := s3.DownloadFile(
		ctx,
		h.s3Client,
		h.cfg.S3RawBucket,
		payload.S3Key,
		inputPath,
	); err != nil {
		return err
	}

	log.Println("Download completed")

	log.Println("Generating HLS...")

	if err := h.generateVariants(
		ctx,
		inputPath,
		outputDir,
	); err != nil {
		return err
	}

	log.Println("HLS generated")

	if err := h.ffmpeg.GenerateMasterPlaylist(outputDir); err != nil {
		return err
	}

	log.Println("Master playlist generated")

	thumbnailPath := filepath.Join(
		workspace,
		"thumbnail.jpg",
	)

	log.Println("Generating thumbnail...")

	if err := h.ffmpeg.GenerateThumbnail(
		ctx,
		inputPath,
		thumbnailPath,
	); err != nil {
		log.Printf("GenerateThumbnail failed: %v", err)
		return err
	}

	log.Println("Thumbnail generated")

	log.Println("Uploading thumbnail...")

	log.Println("Uploading HLS...")

	if err := s3.UploadDirectory(
		ctx,
		h.s3Client,
		h.cfg.S3ProdBucket,
		outputDir,
		filepath.Join(
			s3.ProcessedVideosPrefix,
			payload.VideoID,
		),
	); err != nil {
		return err
	}

	log.Println("HLS uploaded")

	if err := s3.UploadFile(
		ctx,
		h.s3Client,
		h.cfg.S3ProdBucket,
		thumbnailPath,
		filepath.Join(
			s3.ProcessedVideosPrefix,
			payload.VideoID,
			"thumbnail.jpg",
		),
	); err != nil {
		return err
	}

	log.Println("Thumbnail uploaded")

	log.Println("Deleting raw video...")

	if err := s3.DeleteObject(
		ctx,
		h.s3Client,
		h.cfg.S3RawBucket,
		payload.S3Key,
	); err != nil {
		return err
	}

	log.Println("Raw video deleted")

	log.Println("Calling CompleteVideo API...")

	err := h.api.CompleteVideo(
		ctx,
		payload.VideoID,
		api.CompleteVideoRequest{
			HLSURL: filepath.ToSlash(
				filepath.Join(
					s3.ProcessedVideosPrefix,
					payload.VideoID,
					"master.m3u8",
				),
			),
			ThumbnailKey: filepath.ToSlash(
				filepath.Join(
					s3.ProcessedVideosPrefix,
					payload.VideoID,
					"thumbnail.jpg",
				),
			),
		},
	)

	if err != nil {
		log.Printf("CompleteVideo API failed: %v", err)
		return err
	}

	log.Println("CompleteVideo API succeeded")
	log.Println("API updated")

	log.Printf(
		"Video %s processed successfully",
		payload.VideoID,
	)

	return nil
}
