package ffmpeg

import (
	"bytes"
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func (s *Service) GenerateThumbnail(
	ctx context.Context,
	inputPath string,
	outputPath string,
) error {

	if err := os.MkdirAll(
		filepath.Dir(outputPath),
		0755,
	); err != nil {
		return err
	}

	args := []string{
		"-hide_banner",
		"-loglevel", "error",

		"-y",

		"-ss", "00:00:02",

		"-i", inputPath,

		"-frames:v", "1",
		"-update", "1",

		"-q:v", "2",

		outputPath,
	}

	cmd := exec.CommandContext(
		ctx,
		"ffmpeg",
		args...,
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Printf(
			"thumbnail generation failed:\n%s",
			stderr.String(),
		)
		return err
	}

	return nil
}
