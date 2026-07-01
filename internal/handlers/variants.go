package handlers

import (
	"context"
	"log"
	"path/filepath"
	"sync"

	"github.com/Prateet-Github/worker-go/internal/ffmpeg"
)

func (h *VideoHandler) generateVariants(
	ctx context.Context,
	inputPath string,
	outputDir string,
) error {

	var wg sync.WaitGroup

	errCh := make(chan error, len(ffmpeg.Renditions))

	for _, rendition := range ffmpeg.Renditions {
		r := rendition

		wg.Go(func() { // new syntax for WaitGroup in Go 1.25 intead of wg.Add(1) and defer wg.Done() & go func()
			renditionDir := filepath.Join(
				outputDir,
				r.Name,
			)

			log.Printf("Generating %s...", r.Name)

			if err := h.ffmpeg.GenerateVariant(
				ctx,
				inputPath,
				renditionDir,
				r,
			); err != nil {
				errCh <- err
			}
		})
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}
