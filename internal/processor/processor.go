package processor

import (
	"log"

	"github.com/Prateet-Github/worker-go/internal/types"
)

func Process(job types.VideoJob) error {
	log.Println("Starting processing:", job.VideoID)

	return nil
}
