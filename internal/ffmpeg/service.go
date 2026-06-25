package ffmpeg

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateHLS(
	ctx context.Context,
	inputPath string,
	outputDir string,
) error {

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	args := []string{
		"-y",
		"-i", inputPath,

		"-c:v", "libx264",
		"-c:a", "aac",

		"-f", "hls",
		"-hls_time", "6",
		"-hls_playlist_type", "vod",

		"-hls_segment_filename",
		fmt.Sprintf("%s/segment_%%03d.ts", outputDir),

		fmt.Sprintf("%s/master.m3u8", outputDir),
	}

	cmd := exec.CommandContext(
		ctx,
		"ffmpeg",
		args...,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
