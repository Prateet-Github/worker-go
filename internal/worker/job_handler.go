package worker

import (
	"log"

	"github.com/Prateet-Github/worker-go/internal/processor"
	"github.com/Prateet-Github/worker-go/internal/types"
)

func Handle(job types.VideoJob) {
	log.Println("Processing Video:", job.VideoID)

	if err := processor.Process(job); err != nil {

		log.Println("Processing failed:", err)
		return
	}

	// TODO

	// download from S3
	// transcode
	// upload
	// notify API

	log.Println("Finished Video:", job.VideoID)
}
