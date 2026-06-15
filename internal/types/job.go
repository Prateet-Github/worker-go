package types

type VideoJob struct {
	VideoID string `json:"video_id"`
	S3Key   string `json:"s3_key"`
}
