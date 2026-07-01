package ffmpeg

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func (s *Service) GenerateVariant(
	ctx context.Context,
	inputPath string,
	outputDir string,
	rendition Rendition,
) error {

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	// scale := fmt.Sprintf(
	// 	"scale=%d:%d",
	// 	rendition.Width,
	// 	rendition.Height,
	// )

	scale := fmt.Sprintf(
		"scale=-2:%d",
		rendition.Height,
	)

	args := []string{
		"-hide_banner",
		"-loglevel", "error",

		"-y",
		"-i", inputPath,

		"-vf", scale,

		"-c:v", "libx264",
		"-b:v", rendition.VideoBitrate,

		"-c:a", "aac",
		"-b:a", rendition.AudioBitrate,

		"-f", "hls",
		"-hls_time", "6",
		"-hls_playlist_type", "vod",

		"-hls_segment_filename",
		fmt.Sprintf("%s/segment_%%03d.ts", outputDir),

		fmt.Sprintf("%s/index.m3u8", outputDir),
	}

	cmd := exec.CommandContext(
		ctx,
		"ffmpeg",
		args...,
	)

	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// return cmd.Run()

	var stderr bytes.Buffer

	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Printf(
			"ffmpeg failed for %s:\n%s",
			rendition.Name,
			stderr.String(),
		)

		return err
	}

	return nil

}
