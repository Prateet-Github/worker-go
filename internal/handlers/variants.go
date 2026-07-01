package handlers

import (
	"context"
	"log"
	"path/filepath"

	"github.com/Prateet-Github/worker-go/internal/ffmpeg"
)

func (h *VideoHandler) generateVariants(
	ctx context.Context,
	inputPath string,
	outputDir string,
) error {

	for _, rendition := range ffmpeg.Renditions {

		renditionDir := filepath.Join(
			outputDir,
			rendition.Name,
		)

		log.Printf("Generating %s...", rendition.Name)

		if err := h.ffmpeg.GenerateVariant(
			ctx,
			inputPath,
			renditionDir,
			rendition,
		); err != nil {
			return err
		}
	}

	return nil
}
