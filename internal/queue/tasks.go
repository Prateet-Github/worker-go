package queue

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypeProcessVideo = "video:process"

type VideoTask struct {
	VideoID string `json:"videoId"`
	S3Key   string `json:"s3Key"`
}

func NewProcessVideoTask(payload VideoTask) (*asynq.Task, error) {
	data, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeProcessVideo, data), nil
}
