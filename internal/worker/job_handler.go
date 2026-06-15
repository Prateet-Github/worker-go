package worker

import (
	"log"

	"github.com/Prateet-Github/worker-go/internal/types"
)

func Handle(job types.VideoJob) {
	log.Println("Processing Video:", job.VideoID)

	// TODO

	// download from S3
	// transcode
	// upload
	// notify API

	log.Println("Finished Video:", job.VideoID)
}
