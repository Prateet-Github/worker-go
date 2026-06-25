package ffmpeg

import (
	"context"
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
		"-y",

		"-ss", "00:00:02",

		"-i", inputPath,

		"-frames:v", "1",

		"-q:v", "2",

		outputPath,
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
